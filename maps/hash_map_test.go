package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	m := map[string]any{"b": "b", "a": "a", "d": "d", "c": "c"}
	var hashMap = NewHashMap(map[string]any{"b": "b", "a": "a", "d": "d", "c": "c"})

	assert.True(t, hashMap.Exist("a"))
	assert.Equal(t, len(m), hashMap.Count())

	hashMap.ForEach(func(s string, a any) bool {
		if _, ok := m[s]; !ok {
			t.Fatal("值不存在")
		}

		return false
	})

	for k := range m {
		assert.True(t, hashMap.Exist(k))
	}
}

func TestSynchronize(t *testing.T) {
	mep := NewHashMap[string, string]().Synchronize()
	mep.Set("aaaa", "bbb")
}
