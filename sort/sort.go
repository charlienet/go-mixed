package sort

import (
	"fmt"
	"sort"
	"strings"
)

type mapSorter[T any] struct {
	keys []string
	m    map[string]T
}

func SortByKey[T any](m map[string]T) *mapSorter[T] {
	return &mapSorter[T]{
		m:    m,
		keys: keys(m),
	}
}

func (s *mapSorter[T]) Asc() *mapSorter[T] {
	keys := s.keys
	sort.Strings(keys)

	return &mapSorter[T]{
		m:    s.m,
		keys: keys,
	}
}

func (s *mapSorter[T]) Desc() *mapSorter[T] {
	keys := s.keys
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	return &mapSorter[T]{
		m:    s.m,
		keys: keys,
	}
}

func (s *mapSorter[T]) Join(sep string, f func(k string, v T) string) string {
	slice := make([]string, 0, len(s.m))

	keys := s.keys[:]
	for _, k := range keys {
		slice = append(slice, f(k, s.m[k]))
	}

	return strings.Join(slice, sep)
}

func (s *mapSorter[T]) Keys() []string {
	return s.keys
}

func (s *mapSorter[T]) Values() []T {
	ret := make([]T, 0, len(s.m))

	keys := s.keys[:]
	for _, k := range keys {
		ret = append(ret, s.m[k])
	}

	return ret
}

func (s *mapSorter[T]) String() string {
	return fmt.Sprintf("map[%s]", s.Join(" ", func(k string, v T) string {
		return fmt.Sprintf("%s:%v", k, v)
	}))
}

func keys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
