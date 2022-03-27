package generics_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/charlienet/go-mixed/collections/generics"
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
