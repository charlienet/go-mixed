package maps

import (
	"sync"

	"golang.org/x/exp/constraints"
)

var _ Map[string, any] = &rw_map[string, any]{}

type rw_map[K constraints.Ordered, V any] struct {
	m  Map[K, V]
	mu sync.RWMutex
}

func NewRWMap[K constraints.Ordered, V any]() *rw_map[K, V] {
	return &rw_map[K, V]{}
}

func newRWMap[K constraints.Ordered, V any](m Map[K, V]) *rw_map[K, V] {
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
	m.mu.RLock()
	defer m.mu.RUnlock()

	m.m.Delete(key)
}

func (m *rw_map[K, V]) Exist(key K) bool {
	return m.m.Exist(key)
}

func (m *rw_map[K, V]) Count() int {
	return m.m.Count()
}

func (m *rw_map[K, V]) Iter() <-chan *Entry[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.m.Iter()
}

func (m *rw_map[K, V]) ForEach(f func(K, V)) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	m.m.ForEach(f)
}

func (m *rw_map[K, V]) Clone() Map[K, V] {
	return newRWMap(m.m.Clone())
}

func (m *rw_map[K, V]) Clear() {
	m.mu.Lock()
	m.m.Clear()
	m.mu.Unlock()
}
