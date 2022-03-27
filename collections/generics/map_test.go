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
