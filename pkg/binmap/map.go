// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package binmap

const rootRef = ref(1)

// New ...
func New() (m *Map) {
	m = &Map{
		rootBin: 63,
	}
	m.allocCell()
	return
}

// Map ...
type Map struct {
	freeTop    ref
	allocCount int
	cellCount  int
	rootBin    Bin
	cells      []cell
	history    traceHistory
}

// RootBin ...
func (m *Map) RootBin() Bin {
	return m.rootBin
}

func (m *Map) cell(r ref) (c dataCell, mc mapCell) {
	c = dataCell{&m.cells[r]}
	mc = mapCell{&m.cells[r.MapRef()], r}
	return
}

func (m *Map) dataCell(r ref) dataCell {
	return dataCell{&m.cells[r]}
}

func (m *Map) mapCell(r ref) mapCell {
	return mapCell{&m.cells[r.MapRef()], r}
}

func (m *Map) allocCell() (r ref) {
	m.reserveCells(1)

	r = m.freeTop
	c := freeCell{&m.cells[r]}
	mc := m.mapCell(r)

	m.freeTop = c.NextRef()
	m.allocCount++

	c.Reset()
	mc.Reset()
	return
}

func (m *Map) freeCell(r ref) {
	c, mc := m.cell(r)
	if mc.LeftRef() {
		m.freeCell(c.LeftRef())
	}
	if mc.RightRef() {
		m.freeCell(c.RightRef())
	}

	freeCell{&m.cells[r]}.SetNextRef(m.freeTop)
	m.freeTop = r
	m.allocCount--
}

func max(n int, ns ...int) int {
	for i := range ns {
		if ns[i] > n {
			n = ns[i]
		}
	}
	return n
}

func (m *Map) reserveCells(n int) {
	if m.cellCount-m.allocCount > n+1 {
		return
	}

	l := len(m.cells)
	nl := max(
		l*3/2,         // 1.5 growth factor
		l+(n+1)*32/31, // n + cell map overhead
		16,            // min size
	)

	cells := m.cells
	m.cells = make([]cell, nl)
	copy(m.cells, cells)

	for i := nl - 1; i >= l; i-- {
		if !ref(i).IsMapRef() {
			freeCell{&m.cells[i]}.SetNextRef(m.freeTop)
			m.freeTop = ref(i)
			m.cellCount++
		}
	}
}

func (m *Map) extendRoot() {
	c, mc := m.cell(rootRef)
	if !mc.HasRef() && c.Symmetric() {
		c.ResetRight()
	} else {
		r := m.allocCell()

		c, mc := m.cell(rootRef)
		nc, nmc := m.cell(r)
		nc.Copy(c)
		nmc.Copy(mc)

		c.SetLeftRef(r)
		c.ResetRight()
		mc.SetLeftRef(true)
		mc.SetRightRef(false)
	}

	m.rootBin = m.rootBin.Parent()
}

func (m *Map) packCells() {
	i := m.history.len - 1
	r := m.history.refs[i]
	if r == rootRef {
		return
	}

	c, mc := m.cell(r)
	if mc.HasRef() || !c.Symmetric() {
		return
	}

	b := c.LeftBitmap()
	for {
		i--
		r = m.history.refs[i]
		c, mc = m.cell(r)

		if !mc.LeftRef() {
			if c.LeftBitmap() != b {
				break
			}
		} else if !mc.RightRef() {
			if c.RightBitmap() != b {
				break
			}
		} else {
			break
		}

		if r == rootRef {
			break
		}
	}

	nr := m.history.refs[i+1]
	if mc.LeftRef() && c.LeftRef() == nr {
		mc.SetLeftRef(false)
		c.SetLeftBitmap(b)
	} else {
		mc.SetRightRef(false)
		c.SetRightBitmap(b)
	}
	m.freeCell(nr)
}

func (m *Map) trace(target Bin) (r ref, b Bin) {
	r = rootRef
	b = m.rootBin

	for target != b {
		if target < b {
			if m.mapCell(r).LeftRef() {
				r = m.dataCell(r).LeftRef()
				b = b.Left()
			} else {
				break
			}
		} else {
			if m.mapCell(r).RightRef() {
				r = m.dataCell(r).RightRef()
				b = b.Right()
			} else {
				break
			}
		}
	}
	return
}

func (m *Map) traceHistory(target Bin) (r ref, b Bin) {
	r = rootRef
	b = m.rootBin

	m.history.Reset()
	m.history.Append(r)
	for target != b {
		c, mc := m.cell(r)
		if target < b {
			if mc.LeftRef() {
				r = c.LeftRef()
				b = b.Left()
			} else {
				break
			}
		} else {
			if mc.RightRef() {
				r = c.RightRef()
				b = b.Right()
			} else {
				break
			}
		}
		m.history.Append(r)
	}

	return
}

// Set ...
func (m *Map) Set(b Bin) {
	if b.IsNone() {
		return
	}
	if b.LayerBits() > bitmapLayerBits {
		m.setHighLayerBitmap(b, bitmapFilled)
	} else {
		m.setLowLayerBitmap(b, bitmapFilled)
	}
}

// Reset ...
func (m *Map) Reset(b Bin) {
	if b.LayerBits() > bitmapLayerBits {
		m.setHighLayerBitmap(b, bitmapEmpty)
	} else {
		m.setLowLayerBitmap(b, bitmapEmpty)
	}
}

// FillBefore ...
func (m *Map) FillBefore(b Bin) {
	b = b.LayerLeft()
	for !b.IsNone() {
		if b.IsLeft() {
			m.Set(b)
			b = b.Parent().LayerLeft()
		} else {
			b = b.Parent()
		}
	}
}

// ResetBefore ...
func (m *Map) ResetBefore(b Bin) {
	b = b.LayerLeft()
	for !b.IsNone() {
		if b.IsLeft() {
			m.Reset(b)
			b = b.Parent().LayerLeft()
		} else {
			b = b.Parent()
		}
	}
}

func (m *Map) setLowLayerBitmap(target Bin, _bitmap bitmap) {
	binBitmap := binBitmaps[target&bitmapLayerBits]
	bitmap := _bitmap & binBitmap

	if !m.rootBin.Contains(target) {
		if bitmap.Empty() {
			return
		}
		for {
			m.extendRoot()
			if m.rootBin.Contains(target) {
				break
			}
		}
	}

	preBin := (target & ^(bitmapLayerBits + 1)) | bitmapLayerBits
	r, b := m.traceHistory(target)

	c := m.dataCell(r)
	bm := bitmapEmpty
	if target < b {
		bm = c.LeftBitmap()
		if bm&binBitmap == bitmap {
			return
		}
		if b == preBin {
			c.SetLeftBitmap(c.LeftBitmap()&^binBitmap | bitmap)
			m.packCells()
			return
		}
	} else {
		bm = c.RightBitmap()
		if bm&binBitmap == bitmap {
			return
		}
		if b == preBin {
			c.SetRightBitmap(c.RightBitmap()&^binBitmap | bitmap)
			m.packCells()
			return
		}
	}

	m.reserveCells(int(b.Layer() - preBin.Layer()))

	for {
		nr := m.allocCell()
		m.dataCell(nr).SetBitmap(bm)

		c, mc := m.cell(r)
		if preBin < b {
			c.SetLeftRef(nr)
			mc.SetLeftRef(true)
			b = b.Left()
		} else {
			c.SetRightRef(nr)
			mc.SetRightRef(true)
			b = b.Right()
		}

		r = nr

		if b == preBin {
			break
		}
	}

	c = m.dataCell(r)
	if target < b {
		c.SetLeftBitmap(c.LeftBitmap()&^binBitmap | bitmap)
	} else {
		c.SetRightBitmap(c.RightBitmap()&^binBitmap | bitmap)
	}
}

func (m *Map) setHighLayerBitmap(target Bin, _bitmap bitmap) {
	if target.Contains(m.rootBin) {
		c, mc := m.cell(rootRef)
		if mc.LeftRef() {
			m.freeCell(c.LeftRef())
		}
		if mc.RightRef() {
			m.freeCell(c.RightRef())
		}

		m.rootBin = target
		mc.Reset()
		c.SetLeftBitmap(_bitmap)
		c.SetRightBitmap(_bitmap)

		return
	}

	preBin := target.Parent()

	if !m.rootBin.Contains(preBin) {
		if _bitmap.Empty() {
			return
		}

		for {
			m.extendRoot()
			if m.rootBin.Contains(preBin) {
				break
			}
		}
	}

	r, b := m.traceHistory(preBin)
	c, mc := m.cell(r)

	var bm bitmap
	if target < b {
		if mc.LeftRef() {
			mc.SetLeftRef(false)
			m.freeCell(c.LeftRef())
		} else {
			bm = c.LeftBitmap()
			if bm == _bitmap {
				return
			}
		}
		if b == preBin {
			c.SetLeftBitmap(_bitmap)
			m.packCells()
			return
		}
	} else {
		if mc.RightRef() {
			mc.SetRightRef(false)
			m.freeCell(c.RightRef())
		} else {
			bm = c.RightBitmap()
			if bm == _bitmap {
				return
			}
		}
		if b == preBin {
			c.SetRightBitmap(_bitmap)
			m.packCells()
			return
		}
	}

	m.reserveCells(int(b.Layer() - preBin.Layer()))

	for {
		nr := m.allocCell()
		m.dataCell(nr).SetBitmap(bm)

		c, mc := m.cell(r)
		if preBin < b {
			c.SetLeftRef(nr)
			mc.SetLeftRef(true)
			b = b.Left()
		} else {
			c.SetRightRef(nr)
			mc.SetRightRef(true)
			b = b.Right()
		}

		r = nr

		if b == preBin {
			break
		}
	}

	c = m.dataCell(r)
	if target < b {
		c.SetLeftBitmap(_bitmap)
	} else {
		c.SetRightBitmap(_bitmap)
	}
}

func (m *Map) testBin(target Bin, v bitmap) bool {
	r, b := m.trace(target)
	c := m.dataCell(r)

	if target.LayerBits() > bitmapLayerBits {
		if target < b {
			return c.LeftBitmap() == v
		}
		if target > b {
			return c.RightBitmap() == v
		}
		return !m.mapCell(r).HasRef() && c.LeftBitmap() == v && c.RightBitmap() == v
	}

	var bm bitmap
	if target < b {
		bm = c.LeftBitmap()
	} else {
		bm = c.RightBitmap()
	}
	mask := binBitmaps[bitmapLayerBits&target]
	return (bm & mask) == (v & mask)
}

// Empty ...
func (m *Map) Empty() bool {
	c, mc := m.cell(rootRef)
	return !mc.HasRef() && c.LeftBitmap().Empty() && c.RightBitmap().Empty()
}

// Filled ...
func (m *Map) Filled() bool {
	c, mc := m.cell(rootRef)
	return !mc.HasRef() && c.LeftBitmap().Filled() && c.RightBitmap().Filled()
}

// EmptyAt ...
func (m *Map) EmptyAt(b Bin) bool {
	if !m.rootBin.Contains(b) {
		return !b.Contains(m.rootBin) || m.Empty()
	}
	return m.testBin(b, bitmapEmpty)
}

// FilledAt ...
func (m *Map) FilledAt(b Bin) bool {
	if !m.rootBin.Contains(b) {
		return false
	}
	return m.testBin(b, bitmapFilled)
}

// Cover ...
func (m *Map) Cover(target Bin) Bin {
	if !m.rootBin.Contains(target) {
		if !target.Contains(m.rootBin) {
			return m.rootBin.Sibling()
		}
		if m.Empty() {
			return All
		}
		return None
	}

	r, b := m.trace(target)
	c := m.dataCell(r)

	if target.LayerBits() > bitmapLayerBits {
		if target < b {
			if c.LeftBitmap().Empty() || c.LeftBitmap().Filled() {
				return b.Left()
			}
			return None
		}
		if b < target {
			if c.RightBitmap().Empty() || c.RightBitmap().Filled() {
				return b.Right()
			}
			return None
		}
		mc := m.mapCell(r)
		if mc.LeftRef() || mc.RightRef() {
			return None
		}
		if !c.Symmetric() {
			return None
		}
		if c.LeftBitmap().Empty() {
			return All
		}
		if c.LeftBitmap().Filled() {
			return b
		}
		return None
	}

	var bm bitmap
	if target < b {
		bm = c.LeftBitmap()
		b = b.Left()
	} else {
		bm = c.RightBitmap()
		b = b.Right()
	}

	if bm.Empty() {
		if m.Empty() {
			return All
		}
		return b
	}
	if bm.Filled() {
		if m.Filled() {
			return All
		}
		return b
	}

	nb := target
	binBitmap := binBitmaps[nb&bitmapLayerBits]

	if (bm & binBitmap).Empty() {
		for {
			b = nb
			nb = nb.Parent()
			binBitmap = binBitmaps[nb&bitmapLayerBits]

			if !(bm & binBitmap).Empty() {
				return b
			}
		}
	} else if (bm & binBitmap) == binBitmap {
		for {
			b = nb
			nb = nb.Parent()
			binBitmap = binBitmaps[nb&bitmapLayerBits]

			if (bm & binBitmap) != binBitmap {
				return b
			}
		}
	}

	return None
}

// FindEmpty ...
func (m *Map) FindEmpty() Bin {
	var r ref
	var b Bin

	c, mc := m.cell(rootRef)
	bm := bitmapFilled
	if mc.LeftRef() {
		r = c.LeftRef()
		b = m.rootBin.Left()
	} else if !c.LeftBitmap().Filled() {
		if c.LeftBitmap().Empty() {
			if !mc.RightRef() && c.RightBitmap().Empty() {
				return All
			}
			return m.rootBin.Left()
		}
		b = m.rootBin.Left()
		bm = c.LeftBitmap()
		return offsetBitmapBin(b, ^bm)
	} else if mc.RightRef() {
		r = c.RightRef()
		b = m.rootBin.Right()
	} else {
		if c.RightBitmap().Filled() {
			if m.rootBin == All {
				return None
			}
			return m.rootBin.Sibling()
		}
		b = m.rootBin.Right()
		bm = c.RightBitmap()
		return offsetBitmapBin(b, ^bm)
	}

	for {
		c, mc := m.cell(r)
		if mc.LeftRef() {
			r = c.LeftRef()
			b = b.Left()
		} else if !c.LeftBitmap().Filled() {
			bm = c.LeftBitmap()
			b = b.Left()
			break
		} else if mc.RightRef() {
			r = c.RightRef()
			b = b.Right()
		} else {
			bm = c.RightBitmap()
			b = b.Right()
			break
		}
	}

	return offsetBitmapBin(b, ^bm)
}

// Clone ...
func (m *Map) Clone() *Map {
	c := &Map{
		freeTop:    m.freeTop,
		allocCount: m.allocCount,
		cellCount:  m.cellCount,
		rootBin:    m.rootBin,
		cells:      make([]cell, len(m.cells)),
	}
	copy(c.cells, m.cells)
	return c
}

func (m *Map) FindEmptyAfter(target Bin) Bin {
	if !m.rootBin.Contains(target) {
		if target.Contains(m.rootBin) {
			return m.FindEmpty()
		}
		return target
	}

	if !target.IsBase() {
		target = target.BaseLeft()
	}

	r, b := m.traceHistory(target)
	c, mc := m.cell(r)
	var bm bitmap

	mask := binBitmaps[bitmapLayerBits&target] - 1
	bmBaseLeft := target &^ bitmapLayerBits
	if target < b {
		if !(c.LeftBitmap() | mask).Filled() {
			b = bmBaseLeft
			bm = c.LeftBitmap()
			return offsetBitmapBin(b, ^(bm | mask)).BaseLeft()
		} else if b.Left().BaseRight()&^bitmapLayerBits != bmBaseLeft && !c.LeftBitmap().Filled() {
			b = bmBaseLeft + bitmapLayerBits + 1
			bm = c.LeftBitmap()
			return offsetBitmapBin(b, ^bm).BaseLeft()
		} else if mc.RightRef() {
			b = b.Right()
			r = c.RightRef()
			goto Descend
		} else if !c.RightBitmap().Filled() {
			b = b.Right()
			bm = c.RightBitmap()
			return offsetBitmapBin(b, ^bm).BaseLeft()
		}
	} else {
		if !(c.RightBitmap() | mask).Filled() {
			b = bmBaseLeft
			bm = c.RightBitmap()
			return offsetBitmapBin(b, ^(bm | mask)).BaseLeft()
		} else if b.BaseRight()&^bitmapLayerBits != bmBaseLeft && !c.RightBitmap().Filled() {
			b = bmBaseLeft + bitmapLayerBits + 1
			bm = c.RightBitmap()
			return offsetBitmapBin(b, ^bm).BaseLeft()
		}
	}

	for i := m.history.len - 2; ; i-- {
		if b == m.rootBin {
			return m.rootBin.Sibling().BaseLeft()
		}

		b = b.Parent()
		c, mc = m.cell(m.history.refs[i])
		if !mc.RightRef() {
			if !c.RightBitmap().Filled() {
				b = b.Right()
				bm = c.RightBitmap()
				return offsetBitmapBin(b, ^bm).BaseLeft()
			}
		} else if c.RightRef() != r && b > target {
			b = b.Right()
			r = c.RightRef()
			break
		}
	}

Descend:
	for {
		c, mc := m.cell(r)
		if mc.LeftRef() {
			b = b.Left()
			r = c.LeftRef()
		} else if !c.LeftBitmap().Filled() {
			b = b.Left()
			bm = c.LeftBitmap()
			break
		} else if mc.RightRef() {
			b = b.Right()
			r = c.RightRef()
		} else {
			b = b.Right()
			bm = c.RightBitmap()
			break
		}
	}

	return offsetBitmapBin(b, ^bm).BaseLeft()
}

// FindFilled ...
func (m *Map) FindFilled() Bin {
	var r ref
	var b Bin
	c, mc := m.cell(rootRef)
	bm := bitmapEmpty
	if mc.LeftRef() {
		r = c.LeftRef()
		b = m.rootBin.Left()
	} else if !c.LeftBitmap().Empty() {
		if c.LeftBitmap().Filled() {
			if !mc.RightRef() && c.RightBitmap().Filled() {
				return m.rootBin
			}
			return m.rootBin.Left()
		}
		b = m.rootBin.Left()
		bm = c.LeftBitmap()
		return offsetBitmapBin(b, bm)
	} else if mc.RightRef() {
		r = c.RightRef()
		b = m.rootBin.Right()
	} else {
		if c.RightBitmap().Empty() {
			return None
		}
		b = m.rootBin.Right()
		bm = c.RightBitmap()
		return offsetBitmapBin(b, bm)
	}

	for {
		c, mc := m.cell(r)
		if mc.LeftRef() {
			r = c.LeftRef()
			b = b.Left()
		} else if !c.LeftBitmap().Empty() {
			bm = c.LeftBitmap()
			b = b.Left()
			break
		} else if mc.RightRef() {
			r = c.RightRef()
			b = b.Right()
		} else {
			bm = c.RightBitmap()
			b = b.Right()
			break
		}
	}

	return offsetBitmapBin(b, bm)
}

func (m *Map) FindFilledAfter(target Bin) Bin {
	if !m.rootBin.Contains(target) {
		return None
	}

	if !target.IsBase() {
		target = target.BaseLeft()
	}

	r, b := m.traceHistory(target)
	c, mc := m.cell(r)
	var bm bitmap

	mask := binBitmaps[bitmapLayerBits&target] - 1
	bmBaseLeft := target &^ bitmapLayerBits
	if target < b {
		if !(c.LeftBitmap() &^ mask).Empty() {
			b = bmBaseLeft
			bm = c.LeftBitmap()
			return offsetBitmapBin(b, bm&^mask).BaseLeft()
		} else if b.Left().BaseRight()&^bitmapLayerBits != bmBaseLeft && !c.LeftBitmap().Empty() {
			b = bmBaseLeft + bitmapLayerBits + 1
			bm = c.LeftBitmap()
			return offsetBitmapBin(b, bm).BaseLeft()
		} else if mc.RightRef() {
			b = b.Right()
			r = c.RightRef()
			goto Descend
		} else if !c.RightBitmap().Empty() {
			b = b.Right()
			bm = c.RightBitmap()
			return offsetBitmapBin(b, bm).BaseLeft()
		}
	} else {
		if !(c.RightBitmap() &^ mask).Empty() {
			b = bmBaseLeft
			bm = c.RightBitmap()
			return offsetBitmapBin(b, bm&^mask).BaseLeft()
		} else if b.BaseRight()&^bitmapLayerBits != bmBaseLeft && !c.RightBitmap().Empty() {
			b = bmBaseLeft + bitmapLayerBits + 1
			bm = c.RightBitmap()
			return offsetBitmapBin(b, bm).BaseLeft()
		}
	}

	for i := m.history.len - 2; ; i-- {
		if b == m.rootBin {
			return None
		}

		b = b.Parent()
		c, mc = m.cell(m.history.refs[i])
		if !mc.RightRef() {
			if !c.RightBitmap().Empty() {
				b = b.Right()
				bm = c.RightBitmap()
				return offsetBitmapBin(b, bm).BaseLeft()
			}
		} else if c.RightRef() != r && b > target {
			b = b.Right()
			r = c.RightRef()
			break
		}
	}

Descend:
	for {
		c, mc := m.cell(r)
		if mc.LeftRef() {
			b = b.Left()
			r = c.LeftRef()
		} else if !c.LeftBitmap().Empty() {
			b = b.Left()
			bm = c.LeftBitmap()
			break
		} else if mc.RightRef() {
			b = b.Right()
			r = c.RightRef()
		} else {
			b = b.Right()
			bm = c.RightBitmap()
			break
		}
	}

	return offsetBitmapBin(b, bm).BaseLeft()
}

// FindLastFilled ...
func (m *Map) FindLastFilled() Bin {
	if m.Empty() {
		return None
	}

	b := m.rootBin
	for b.Layer() != 0 {
		if m.EmptyAt(b.Right()) {
			b = b.Left()
		} else {
			b = b.Right()
		}
	}
	return b
}

// IterateEmpty ...
func (m *Map) IterateEmpty() Iterator {
	return NewEmptyAtIterator(m, m.rootBin)
}

// IterateFilled ...
func (m *Map) IterateFilled() Iterator {
	return NewFilledAtIterator(m, m.rootBin)
}

// IterateEmptyAt ...
func (m *Map) IterateEmptyAt(b Bin) Iterator {
	return NewEmptyAtIterator(m, b)
}

// IterateFilledAt ...
func (m *Map) IterateFilledAt(b Bin) Iterator {
	return NewFilledAtIterator(m, b)
}

type traceHistory struct {
	refs [64]ref
	len  int
}

func (t *traceHistory) Append(r ref) {
	t.refs[t.len] = r
	t.len++
}

func (t *traceHistory) Slice() (r []ref) {
	r = make([]ref, t.len)
	for i := 0; i < t.len; i++ {
		r[i] = t.refs[i]
	}
	return
}

func (t *traceHistory) Reset() {
	t.len = 0
}
