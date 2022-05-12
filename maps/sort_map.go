package maps

import (
	"fmt"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

var (
	_ Map[string, any]       = &sorted_map[string, any]{}
	_ SortedMap[string, any] = &sorted_map[string, any]{}
)

type SortedMap[K constraints.Ordered, V any] interface {
	Map[K, V]
	Asc() SortedMap[K, V]
	Desc() SortedMap[K, V]
}

type sorted_map[K constraints.Ordered, V any] struct {
	keys []K
	maps Map[K, V]
}

func NewSortedMap[K constraints.Ordered, V any](maps ...map[K]V) *sorted_map[K, V] {
	merged := Merge(maps...)
	return &sorted_map[K, V]{
		keys: keys(merged),
		maps: NewHashMap(merged),
	}
}

func NewSortedByMap[K constraints.Ordered, V any](m Map[K, V]) *sorted_map[K, V] {
	return &sorted_map[K, V]{maps: m, keys: m.Keys()}
}

func (m *sorted_map[K, V]) Get(key K) (V, bool) {
	return m.maps.Get(key)
}

func (m *sorted_map[K, V]) Set(key K, value V) {
	m.maps.Set(key, value)
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

func (m *sorted_map[K, V]) Clear() {
	m.keys = make([]K, 0)
	m.maps.Clear()
}

func (m *sorted_map[K, V]) Clone() Map[K, V] {
	return &sorted_map[K, V]{maps: m.maps.Clone(), keys: m.Keys()}
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

func (m *sorted_map[K, V]) ForEach(f func(K, V)) {
	for _, k := range m.keys {
		if v, ok := m.Get(k); ok {
			f(k, v)
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

func (m *sorted_map[K, V]) String() string {
	return fmt.Sprintf("map[%s]", Join[K, V](m, " ", func(k K, v V) string {
		return fmt.Sprintf("%v:%v", k, v)
	}))
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

	slices.SortFunc(keys, func(a, b K) bool {
		return a > b
	})

	return &sorted_map[K, V]{
		maps: m.maps,
		keys: keys,
	}
}

func keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
