package maps

import (
	"github.com/alphadose/haxmap"
)

type haxHashable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string | complex64 | complex128
}

var _ Map[string, string] = &haxmapWrapper[string, string]{}

type haxmapWrapper[K haxHashable, V any] struct {
	mep *haxmap.Map[K, V]
}

func NewHaxmap[K haxHashable, V any](size int) *haxmapWrapper[K, V] {
	return &haxmapWrapper[K, V]{
		mep: haxmap.New[K, V](uintptr(size)),
	}
}

func (m *haxmapWrapper[K, V]) Set(k K, v V) {
	m.mep.Set(k, v)
}

func (m *haxmapWrapper[K, V]) Get(k K) (V, bool) {
	return m.mep.Get(k)
}

func (m *haxmapWrapper[K, V]) Keys() []K {
	keys := make([]K, 0, m.mep.Len())
	m.mep.ForEach(func(k K, v V) bool {
		keys = append(keys, k)

		return true
	})

	return keys
}

func (m *haxmapWrapper[K, V]) Values() []V {
	values := make([]V, 0, m.mep.Len())
	m.mep.ForEach(func(k K, v V) bool {
		values = append(values, v)

		return true
	})

	return values
}

func (m *haxmapWrapper[K, V]) Exist(k K) bool {
	return false
}

func (m *haxmapWrapper[K, V]) Delete(key K) {
	m.mep.Del(key)
}

func (m *haxmapWrapper[K, V]) ToMap() map[K]V {
	mm := make(map[K]V, m.mep.Len())
	m.mep.ForEach(func(k K, v V) bool {
		mm[k] = v

		return true
	})

	return mm
}

func (m *haxmapWrapper[K, V]) ForEach(fn func(K, V) bool) {

}

func (m *haxmapWrapper[K, V]) Count() int {
	return int(m.mep.Len())
}

func (m *haxmapWrapper[K, V]) Clear() {
}
