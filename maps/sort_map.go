package maps

import (
	"fmt"

	"slices"

	"github.com/charlienet/go-mixed/expr"
	xmaps "golang.org/x/exp/maps"
)

var (
	_ SortedMap[string, any] = &sorted_map[string, any]{}
)

type SortedMap[K hashable, V any] interface {
	Map[K, V]
	Asc() SortedMap[K, V]
	Desc() SortedMap[K, V]
}

type sorted_map[K hashable, V any] struct {
	keys []K
	maps Map[K, V]
}

func NewSortedMap[K hashable, V any](maps ...map[K]V) *sorted_map[K, V] {
	merged := Merge(maps...)
	return &sorted_map[K, V]{
		keys: xmaps.Keys(merged),
		maps: NewHashMap(merged),
	}
}

func NewSortedByMap[K hashable, V any](m Map[K, V]) *sorted_map[K, V] {
	return &sorted_map[K, V]{maps: m, keys: m.Keys()}
}

func (m *sorted_map[K, V]) Get(key K) (V, bool) {
	return m.maps.Get(key)
}

func (m *sorted_map[K, V]) Set(key K, value V) {
	m.maps.Set(key, value)

	slices.Sort(m.keys)
	m.keys = append(m.keys, key)
}

func (m *sorted_map[K, V]) Delete(key K) {
	m.maps.Delete(key)

	for index := range m.keys {
		if m.keys[index] == key {
			m.keys = append(m.keys[:index], m.keys[index+1:]...)
			break
		}
	}
}

func (m *sorted_map[K, V]) Count() int {
	return m.maps.Count()
}

func (m *sorted_map[K, V]) Iter() <-chan *Entry[K, V] {
	c := make(chan *Entry[K, V], m.Count())
	go func() {
		for _, k := range m.keys {
			v, _ := m.maps.Get(k)

			c <- &Entry[K, V]{
				Key:   k,
				Value: v,
			}
		}
		close(c)
	}()

	return c
}

func (m *sorted_map[K, V]) ForEach(f func(K, V) bool) {
	keys := m.keys[:]
	for _, k := range keys {
		if v, ok := m.Get(k); ok {
			if f(k, v) {
				break
			}
		}
	}
}

func (m *sorted_map[K, V]) Exist(key K) bool {
	return m.maps.Exist(key)
}

func (m *sorted_map[K, V]) Keys() []K {
	return m.keys
}

func (s *sorted_map[K, V]) Values() []V {
	ret := make([]V, 0, s.maps.Count())
	for _, k := range s.keys {
		v, _ := s.maps.Get(k)
		ret = append(ret, v)
	}

	return ret
}

func (s *sorted_map[K, V]) ToMap() map[K]V {
	return s.maps.ToMap()
}

func (m *sorted_map[K, V]) Asc() SortedMap[K, V] {
	keys := m.keys
	slices.Sort(keys)

	return &sorted_map[K, V]{
		maps: m.maps,
		keys: keys,
	}
}

func (m *sorted_map[K, V]) Desc() SortedMap[K, V] {
	keys := m.keys

	slices.SortFunc(keys, func(a, b K) int {
		if a == b {
			return 0
		}

		return expr.Ternary(a > b, -1, 1)
	})

	return &sorted_map[K, V]{
		maps: m.maps,
		keys: keys,
	}
}

func (m *sorted_map[K, V]) String() string {
	return fmt.Sprintf("map[%s]", Join[K, V](m, " ", func(k K, v V) string {
		return fmt.Sprintf("%v:%v", k, v)
	}))
}
