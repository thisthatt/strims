// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package set

func New[T comparable](size int) *Set[T] {
	return &Set[T]{
		values: make(map[T]struct{}, size),
	}
}

func NewFromSlice[T comparable](vs []T) *Set[T] {
	s := &Set[T]{
		values: make(map[T]struct{}, len(vs)),
	}
	for _, v := range vs {
		s.values[v] = struct{}{}
	}
	return s
}

type Set[T comparable] struct {
	values map[T]struct{}
}

func (s *Set[T]) Insert(v T) {
	if _, ok := s.values[v]; !ok {
		s.values[v] = struct{}{}
	}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.values[v]
	return ok
}

func (s *Set[T]) Slice() []T {
	vs := make([]T, 0, len(s.values))
	for v := range s.values {
		vs = append(vs, v)
	}
	return vs
}
