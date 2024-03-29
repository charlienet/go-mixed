package maps

import (
	"sync"
)

var _ Map[string, any] = &rw_map[string, any]{}

type rw_map[K hashable, V any] struct {
	m  Map[K, V]
	mu sync.RWMutex
}

func NewRWMap[K hashable, V any](maps ...map[K]V) *rw_map[K, V] {
	merged := Merge(maps...)
	return &rw_map[K, V]{m: NewHashMap(merged)}
}

func newRWMap[K hashable, V any](m Map[K, V]) *rw_map[K, V] {
	return &rw_map[K, V]{m: m}
}

func (m *rw_map[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.m.Set(key, value)
	m.mu.Unlock()
}

func (m *rw_map[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.m.Get(key)
}

func (m *rw_map[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m.Delete(key)
}

func (m *rw_map[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.m.Keys()
}

func (m *rw_map[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.m.Values()
}

func (m *rw_map[K, V]) ToMap() map[K]V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.m.ToMap()
}

func (m *rw_map[K, V]) Shrink() map[K]V {
	return m.m.ToMap()
}

func (m *rw_map[K, V]) Exist(key K) bool {
	return m.m.Exist(key)
}

func (m *rw_map[K, V]) Count() int {
	return m.m.Count()
}

// func (m *rw_map[K, V]) Iter() <-chan *Entry[K, V] {
// 	m.mu.RLock()
// 	defer m.mu.RUnlock()

// 	return m.m.Iter()
// }

func (m *rw_map[K, V]) ForEach(f func(K, V) bool) {

}
