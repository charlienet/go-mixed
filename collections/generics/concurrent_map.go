package generics

import (
	"fmt"
	"runtime"
	"sync"
)

var _ Map[string, string] = &ConcurrnetMap[string, string]{}

var defaultNumOfBuckets = runtime.GOMAXPROCS(runtime.NumCPU())

type ConcurrnetMap[K comparable, V any] struct {
	buckets      []Map[K, V]
	numOfBuckets uint64
}

func NewConcurrnetMap[K comparable, V any]() Map[K, V] {
	num := defaultNumOfBuckets

	buckets := make([]Map[K, V], num)
	for i := 0; i < num; i++ {
		buckets[i] = NewRWLockMap[K, V]()
	}

	return &ConcurrnetMap[K, V]{
		numOfBuckets: uint64(num),
		buckets:      buckets,
	}
}

func (m *ConcurrnetMap[K, V]) Set(key K, value V) {
	m.getBucket(key).Set(key, value)
}

func (m *ConcurrnetMap[K, V]) Get(key K) (V, bool) {
	return m.getBucket(key).Get(key)
}

func (m *ConcurrnetMap[K, V]) Delete(key K) {
	im := m.getBucket(key)
	im.Delete(key)
}

func (m *ConcurrnetMap[K, V]) ForEach(f func(K, V)) {
	var wg sync.WaitGroup

	num := int(m.numOfBuckets)

	wg.Add(int(m.numOfBuckets))
	for i := 0; i < num; i++ {
		go func(i int) {
			m.buckets[i].ForEach(f)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func (m *ConcurrnetMap[K, V]) Clone() Map[K, V] {

	num := int(m.numOfBuckets)

	buckets := make([]Map[K, V], m.numOfBuckets)
	for i := 0; i < num; i++ {
		buckets[i] = m.buckets[i].Clone()
	}

	return &ConcurrnetMap[K, V]{
		buckets:      buckets,
		numOfBuckets: m.numOfBuckets,
	}
}

func (m *ConcurrnetMap[K, V]) Clear() {
	for i := 0; i < int(m.numOfBuckets); i++ {
		m.buckets[i].Clear()
	}
}

func (m *ConcurrnetMap[K, V]) Count() int {
	var count int
	for i := 0; i < int(m.numOfBuckets); i++ {
		count += m.buckets[i].Count()
	}

	return count
}

func (m *ConcurrnetMap[K, V]) getBucket(k K) Map[K, V] {
	id := getTag(k) % m.numOfBuckets
	return m.buckets[id]
}

func getTag[T comparable](v T) uint64 {
	var vv any = v

	switch vv.(type) {
	case string:
		return fnv64(vv.(string))
	case int8:
		return uint64(vv.(int8))
	case uint8:
		return uint64(vv.(uint8))
	case int:
		return uint64(vv.(int))
	case int32:
		return uint64(vv.(int32))
	case uint32:
		return uint64(vv.(uint32))
	case int64:
		return uint64(vv.(int64))
	case uint64:
		return vv.(uint64)
	default:
		return fnv64(fmt.Sprintf("%v", v))
	}
}

const (
	prime32 = uint64(16777619)
)

func fnv64(k string) uint64 {
	var hash = uint64(2166136261)
	l := len(k)
	for i := 0; i < l; i++ {
		hash *= prime32
		hash ^= uint64(k[i])
	}

	return hash
}
