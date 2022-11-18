package maps

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/hash"
)

var defaultNumOfBuckets = runtime.GOMAXPROCS(runtime.NumCPU())

type concurrnetMap[K hashable, V any] struct {
	buckets      []Map[K, V]
	numOfBuckets uint64
}

func NewConcurrentMap[K hashable, V any](maps ...map[K]V) *concurrnetMap[K, V] {
	num := defaultNumOfBuckets

	buckets := make([]Map[K, V], num)
	for i := 0; i < num; i++ {
		buckets[i] = NewRWMap[K, V]()
	}

	m := &concurrnetMap[K, V]{
		numOfBuckets: uint64(num),
		buckets:      buckets,
	}

	for k := range maps {
		for k, v := range maps[k] {
			m.Set(k, v)
		}
	}

	return m
}

func (m *concurrnetMap[K, V]) Set(key K, value V) {
	m.getBucket(key).Set(key, value)
}

func (m *concurrnetMap[K, V]) Get(key K) (V, bool) {
	return m.getBucket(key).Get(key)
}

func (m *concurrnetMap[K, V]) Delete(key K) {
	im := m.getBucket(key)
	im.Delete(key)
}

func (m *concurrnetMap[K, V]) Exist(key K) bool {
	mm := m.getBucket(key)
	return mm.Exist(key)
}

// func (m *concurrnetMap[K, V]) Iter() <-chan *Entry[K, V] {
// 	num := int(m.numOfBuckets)
// 	ch := make(chan *Entry[K, V], m.Count())
// 	for i := 0; i < num; i++ {
// 		c := m.buckets[i].Iter()
// 		ch <- <-c
// 	}

// 	return ch
// }

func (m *concurrnetMap[K, V]) Keys() []K {
	keys := make([]K, m.Count())
	for _, b := range m.buckets {
		keys = append(keys, b.Keys()...)
	}

	return keys
}

func (m *concurrnetMap[K, V]) Values() []V {
	values := make([]V, 0, m.Count())
	for _, v := range m.buckets {
		values = append(values, v.Values()...)
	}

	return values
}

func (m *concurrnetMap[K, V]) ToMap() map[K]V {
	mm := make(map[K]V, m.Count())
	for _, v := range m.buckets {
		mm = Merge(mm, v.ToMap())
	}

	return mm
}

func (m *concurrnetMap[K, V]) ForEach(f func(K, V) bool) {
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

func (m *concurrnetMap[K, V]) Count() int {
	var count int
	for i := 0; i < int(m.numOfBuckets); i++ {
		count += m.buckets[i].Count()
	}

	return count
}

func (m *concurrnetMap[K, V]) getBucket(k K) Map[K, V] {
	id := getTag(k) % m.numOfBuckets
	return m.buckets[id]
}

func getTag[T comparable](v T) uint64 {
	var vv any = v

	switch vv.(type) {
	case string:
		return fnv(vv.(string))
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
		return fnv(fmt.Sprintf("%v", v))
	}
}

func fnv(k string) uint64 {
	bytes := bytesconv.StringToBytes(k)
	return uint64(hash.Funv32(bytes))
}
