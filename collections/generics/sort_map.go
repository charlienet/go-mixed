package generics

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type mapSorter[K constraints.Ordered, V any] struct {
	keys []K
	m    map[K]V
}

func NewSortMap[K constraints.Ordered, V any](m map[K]V) *mapSorter[K, V] {
	return &mapSorter[K, V]{
		m:    m,
		keys: keys(m),
	}
}

func (s *mapSorter[K, V]) Set(key K, value V) {
	s.m[key] = value
	s.keys = append(s.keys, key)
}

func (s *mapSorter[K, V]) Get(key K) (V, bool) {
	v, ok := s.m[key]
	return v, ok
}

func (s *mapSorter[K, V]) Delete(key K) {
	delete(s.m, key)

	keys := keys(s.m)
	s.keys = keys
}

func (s *mapSorter[K, V]) Clear() {
	s.m = map[K]V{}
	s.keys = []K{}
}

func (s *mapSorter[K, V]) Count() int {
	return len(s.m)
}

func (s *mapSorter[K, V]) Asc() *mapSorter[K, V] {
	keys := s.keys
	slices.Sort(keys)

	return &mapSorter[K, V]{
		m:    s.m,
		keys: keys,
	}
}

func (s *mapSorter[K, V]) Desc() *mapSorter[K, V] {
	keys := s.keys

	slices.SortFunc(keys, func(a, b K) bool {
		return a > b
	})

	return &mapSorter[K, V]{
		m:    s.m,
		keys: keys,
	}
}

func (s *mapSorter[K, V]) Join(sep string, f func(k K, v V) string) string {
	slice := make([]string, 0, len(s.m))
	for _, k := range s.keys {
		slice = append(slice, f(k, s.m[k]))
	}

	return strings.Join(slice, sep)
}

func (s *mapSorter[K, V]) Keys() []K {
	return s.keys
}

func (s *mapSorter[K, V]) Values() []V {
	ret := make([]V, 0, len(s.m))
	for _, k := range s.keys {
		ret = append(ret, s.m[k])
	}

	return ret
}

func (s *mapSorter[K, V]) String() string {
	return fmt.Sprintf("map[%s]", s.Join(" ", func(k K, v V) string {
		return fmt.Sprintf("%v:%v", k, v)
	}))
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
