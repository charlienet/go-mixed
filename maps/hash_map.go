package maps

import (
	"github.com/charlienet/go-mixed/locker"
)

type hashMap[K hashable, V any] struct {
	m   map[K]V
	opt *options
}

func NewHashMap[K hashable, V any](maps ...map[K]V) *hashMap[K, V] {
	m := make(map[K]V)
	if len(maps) > 0 {
		m = Merge(maps...)
	}

	return &hashMap[K, V]{opt: acquireDefaultOptions(), m: m}
}

// Synchronize
func (m *hashMap[K, V]) Synchronize() *hashMap[K, V] {
	m.opt.mu = locker.NewRWLocker()
	m.opt.hasLocker = true

	return m
}

func (m *hashMap[K, V]) Set(key K, value V) {
	m.opt.mu.Lock()
	m.m[key] = value
	m.opt.mu.Unlock()
}

func (m *hashMap[K, V]) Get(key K) (V, bool) {
	m.opt.mu.RLock()
	v, exist := m.m[key]
	m.opt.mu.RUnlock()
	return v, exist
}

func (m *hashMap[K, V]) Delete(key K) {
	m.opt.mu.Lock()
	defer m.opt.mu.Unlock()

	delete(m.m, key)
}

func (m *hashMap[K, V]) Exist(key K) bool {
	m.opt.mu.RLock()
	defer m.opt.mu.RUnlock()

	_, ok := m.m[key]
	return ok
}

func (m *hashMap[K, V]) Iter() <-chan *Entry[K, V] {
	ch := make(chan *Entry[K, V], m.Count())
	go func() {
		m.opt.mu.RLock()
		defer m.opt.mu.RUnlock()

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

func (m *hashMap[K, V]) ForEach(f func(K, V) bool) {
	cloned := m.ToMap()

	for k, v := range cloned {
		f(k, v)
	}
}

func (m *hashMap[K, V]) Keys() []K {
	m.opt.mu.RLock()
	defer m.opt.mu.RUnlock()

	keys := make([]K, 0, m.Count())
	for k := range m.m {
		keys = append(keys, k)
	}

	return keys
}

func (m *hashMap[K, V]) Values() []V {
	m.opt.mu.RLock()
	defer m.opt.mu.RUnlock()

	values := make([]V, 0, m.Count())
	for _, v := range m.m {
		values = append(values, v)
	}

	return values
}

func (m *hashMap[K, V]) ToMap() map[K]V {
	m.opt.mu.RLock()
	defer m.opt.mu.RUnlock()

	mm := make(map[K]V, m.Count())
	for k, v := range m.m {
		mm[k] = v
	}

	return mm
}

func (m *hashMap[K, V]) Clear() {
	m.opt.mu.Lock()
	defer m.opt.mu.Unlock()

	m.m = make(map[K]V)
}

func (m *hashMap[K, V]) Count() int {
	return len(m.m)
}

func (m *hashMap[K, V]) Clone() Map[K, V] {
	ret := NewHashMap(m.ToMap())

	if m.opt.hasLocker {
		ret = ret.Synchronize()
	}

	return ret
}
