// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package syncutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUninitializedPointer(t *testing.T) {
	var p Pointer[string]

	next := "string"
	prev := p.Swap(&next)
	assert.Nil(t, prev)
	assert.Equal(t, &next, p.Get())
}
