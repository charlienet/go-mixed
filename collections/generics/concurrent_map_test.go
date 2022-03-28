package generics_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/collections/generics"
	"github.com/charlienet/go-mixed/hash"
)

func TestConcurrentMap(t *testing.T) {
	t.Log(runtime.GOMAXPROCS(runtime.NumCPU()))

	key := "abc"
	value := "bcd"

	m := generics.NewConcurrnetMap[string, string]()
	m.Set(key, value)
	v, ok := m.Get(key)
	t.Log("v:", v, ok)

	m.Delete(key)
	v, ok = m.Get(key)
	t.Log("v:", v, ok)
}

func TestForEach(t *testing.T) {

	m := generics.NewConcurrnetMap[string, string]()
	for k := 0; k < 10; k++ {
		key := fmt.Sprintf("abc-%d", k)
		value := fmt.Sprintf("abc-%d", k)
		m.Set(key, value)
	}

	m.ForEach(func(s1, s2 string) {
		t.Log(s1, s2)
	})

	t.Log("finish")
}

func TestExist(t *testing.T) {
	key := "abc"
	value := "bcd"

	m := generics.NewConcurrnetMap[string, string]()
	m.Set(key, value)

	keyv2 := "abc"

	t.Logf("%p %p", &key, &keyv2)
	_, ok := m.Get(keyv2)
	t.Log("ok", ok)
}

func BenchmarkConcurrnetMap(b *testing.B) {
	key := "abc"
	value := "bcd"

	m := generics.NewConcurrnetMap[string, string]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Set(key, value)
			m.Get(key)
			m.Delete(key)
		}
	})
}

func BenchmarkRWLockMap(b *testing.B) {
	key := "abc"
	value := "bcd"

	m := generics.NewRWLockMap[string, string]()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Set(key, value)
			m.Get(key)
			m.Delete(key)
		}
	})
}

func BenchmarkGetIndex(b *testing.B) {

	b.Run("字符串-Sprintf", func(b *testing.B) {
		v := "abc"
		for i := 0; i < b.N; i++ {
			bytes := []byte(fmt.Sprintf("%v", v))
			hash.XXHashUint64(bytes)
		}
	})

	b.Run("字符串", func(b *testing.B) {
		v := "abc"
		for i := 0; i < b.N; i++ {
			getTag(v)
		}
	})

	b.Run("数字", func(b *testing.B) {
		v := 124
		for i := 0; i < b.N; i++ {
			getTag(v)
		}
	})
}

func getTag[T comparable](v T) uint64 {
	var vv any = v

	switch vv.(type) {
	case string:
		return getStringIndex([]byte(vv.(string)))
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
		return getStringIndex([]byte(fmt.Sprintf("%v", v)))
	}
}

func getStringIndex(v []byte) uint64 {
	if len(v) > 4 {
		v = v[:4]
	}

	// return hash.XXHashUint64(v)
	i, _ := bytesconv.BigEndian.BytesToUInt64(v)
	return i
}

func TestFnv64(t *testing.T) {
	t.Log(fnv64("a"))
	t.Log(fnv64("b"))
	t.Log(fnv64("c"))
	t.Log(fnv64("d"))
	t.Log(fnv64("e"))
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
