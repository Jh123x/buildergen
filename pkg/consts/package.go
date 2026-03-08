package consts

import (
	"path"
)

func GetBase(s string) string {
	return path.Base(s)
}

type Set[T comparable] map[T]Empty

func NewSet[T comparable](elems ...T) Set[T] {
	s := make(Set[T], 0)
	for _, e := range elems {
		s.Add(e)
	}
	return s
}

func (s Set[T]) Add(v T) {
	s[v] = Empty{}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) ToList() []T {
	v := make([]T, 0, len(s))
	for k := range s {
		v = append(v, k)
	}
	return v
}
