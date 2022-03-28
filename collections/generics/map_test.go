package generics_test

import (
	"testing"

	"github.com/charlienet/go-mixed/collections/generics"
	"github.com/stretchr/testify/assert"
)

func TestIMap(t *testing.T) {
	k := "abc"
	v := "bcd"

	var m generics.Map[string, string] = generics.NewConcurrnetMap[string, string]()
	m.Set(k, v)
	_, ok := m.Get(k)
	assert.True(t, ok, "不存在")
	t.Log(m.Count())

	m.Delete(k)
	_, ok = m.Get(k)
	assert.False(t, ok, "不存在")

	t.Log(m.Count())
}

func TestMapCount(t *testing.T) {
	mm := make(map[string]string)
	mm["a"] = "b"
	assert.Equal(t, 1, len(mm))
}

func BenchmarkMap(b *testing.B) {
	b.Run("RWLock", func(b *testing.B) {
		m := generics.NewRWLockMap[string, string]()
		doBenchamark(b, m)

	})

	b.Run("ConcurrnetMap", func(b *testing.B) {
		doBenchamark(b, generics.NewConcurrnetMap[string, string]())
	})
}

func doBenchamark(b *testing.B, m generics.Map[string, string]) {
	var k = "abc"
	var v = "bcd"

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Set(k, v)
			m.Get(k)
			m.Get(k)
			m.Get(k)
			m.Delete(k)
		}
	})

}
