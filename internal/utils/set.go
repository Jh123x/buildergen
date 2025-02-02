package utils

import "github.com/Jh123x/buildergen/internal/consts"

type Set[T comparable] map[T]consts.Empty

func NewSet[T comparable]() Set[T] {
	return make(Set[T], 0)
}

func (s Set[T]) Add(v T) {
	s[v] = consts.Empty{}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len(v T) int {
	return len(s)
}

func (s Set[T]) ToList() []T {
	v := make([]T, 0, len(s))
	for k := range s {
		v = append(v, k)
	}

	return v
}
