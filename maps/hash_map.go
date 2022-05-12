package maps

import (
	"golang.org/x/exp/constraints"
)

type hashMap[K constraints.Ordered, V any] struct {
	m map[K]V
}

func NewHashMap[K constraints.Ordered, V any](maps ...map[K]V) *hashMap[K, V] {
	m := make(map[K]V)
	if len(maps) > 0 {
		m = Merge(maps...)
	}

	return &hashMap[K, V]{m: m}
}

// synchronized
func (m *hashMap[K, V]) Synchronized() *hashMap[K, V] {
	return m
}

func (m *hashMap[K, V]) Set(key K, value V) {
	m.m[key] = value
}

func (m *hashMap[K, V]) Get(key K) (V, bool) {
	v, exist := m.m[key]
	return v, exist
}

func (m *hashMap[K, V]) Delete(key K) {
	delete(m.m, key)
}

func (m *hashMap[K, V]) Exist(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *hashMap[K, V]) Iter() <-chan *Entry[K, V] {
	ch := make(chan *Entry[K, V], m.Count())
	go func() {
		for k, v := range m.m {
			ch <- &Entry[K, V]{
				Key:   k,
				Value: v,
			}
		}

		close(ch)
	}()

	return ch
}

func (m *hashMap[K, V]) ForEach(f func(K, V)) {
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *hashMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Count())
	for k := range m.m {
		keys = append(keys, k)
	}

	return keys
}

func (m *hashMap[K, V]) Values() []V {
	values := make([]V, 0, m.Count())
	for _, v := range m.m {
		values = append(values, v)
	}

	return values
}

func (m *hashMap[K, V]) ToMap() map[K]V {
	mm := make(map[K]V, m.Count())
	for k, v := range m.m {
		mm[k] = v
	}

	return mm
}

func (m *hashMap[K, V]) Clear() {
	m.m = make(map[K]V)
}

func (m *hashMap[K, V]) Count() int {
	return len(m.m)
}

func (m *hashMap[K, V]) Clone() Map[K, V] {
	return NewHashMap(m.ToMap())
}
