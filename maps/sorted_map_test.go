package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortMapConvert(t *testing.T) {
	var m = NewSortedMap(map[string]any{"aaa": "bbb"}).Asc()

	var c Map[string, any] = m

	t.Log(Join(c, " ", func(k string, v any) string {
		return fmt.Sprintf("%s=%v", k, v)
	}))
}

func TestSortedJoin(t *testing.T) {
	m := NewSortedMap(map[string]any{"b": "b", "a": "a", "d": "d", "c": "c"}).Asc()

	f := func(k string, v any) string {
		return fmt.Sprintf("%s=%v", k, v)
	}

	ret := Join[string, any](m, "&", f)

	assert.Equal(t, ret, "a=a&b=b&c=c&d=d")

	ret = Join[string, any](m.Desc(), "&", f)

	assert.Equal(t, "d=d&c=c&b=b&a=a", ret)
}
