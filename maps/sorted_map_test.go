package maps

import (
	"fmt"
	"testing"
)

func TestSortMapConvert(t *testing.T) {
	var m = NewSortedMap(map[string]any{"aaa": "bbb"}).Asc()

	var c Map[string, any] = m

	t.Log(Join(c, " ", func(k string, v any) string {
		return fmt.Sprintf("%s=%v", k, v)
	}))
}

func TestSortedJoin(t *testing.T) {
	var m = NewSortedMap(map[string]any{"b": "b", "a": "a", "d": "d", "c": "c"}).Asc()

	t.Log(Join[string, any](m, "&", func(k string, v any) string {
		return fmt.Sprintf("%s=%v", k, v)
	}))
}
