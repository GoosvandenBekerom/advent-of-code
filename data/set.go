package data

import "golang.org/x/exp/maps"

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s Set[T]) Add(i T) {
	s.m[i] = struct{}{}
}

func (s Set[T]) Has(i T) bool {
	_, exists := s.m[i]
	return exists
}

func (s Set[T]) Remove(i T) {
	delete(s.m, i)
}

func (s Set[T]) Len() int {
	return len(s.m)
}

func (s Set[T]) Values() (values []T) {
	return maps.Keys(s.m)
}
