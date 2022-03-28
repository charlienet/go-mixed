package generics

import (
	"sync"

	"golang.org/x/exp/maps"
)

var _ Map[string, string] = &RWLockMap[string, string]{}

type RWLockMap[K comparable, V any] struct {
	m    map[K]V
	lock sync.RWMutex
}

func NewRWLockMap[K comparable, V any]() Map[K, V] {
	return &RWLockMap[K, V]{
		m: make(map[K]V),
	}
}

func (m *RWLockMap[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	v, ok := m.m[key]
	m.lock.RUnlock()
	return v, ok
}

func (m *RWLockMap[K, V]) Set(key K, value V) {
	m.lock.Lock()
	m.m[key] = value
	m.lock.Unlock()
}

func (m *RWLockMap[K, V]) Delete(key K) {
	m.lock.Lock()
	delete(m.m, key)
	m.lock.Unlock()
}

func (m *RWLockMap[K, V]) ForEach(f func(K, V)) {
	for k, v := range m.m {
		f(k, v)
	}
}

func (m *RWLockMap[K, V]) Clone() Map[K, V] {
	new := maps.Clone(m.m)

	return &RWLockMap[K, V]{
		m: new,
	}
}

func (m *RWLockMap[K, V]) Clear() {
	m.m = make(map[K]V)
}

func (m *RWLockMap[K, V]) Count() int {
	return len(m.m)
}
