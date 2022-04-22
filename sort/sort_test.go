package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	m := map[string]any{
		"Bcd": "bcd",
		"Abc": "abc",
	}

	t.Log(SortByKey(m).Values())
	t.Log(SortByKey(m).Desc().Values())
	t.Log(SortByKey(m).Asc().Values())
}

func TestMapSortInt(t *testing.T) {
	m := map[string]int{
		"Bcd": 8,
		"Abc": 4,
	}

	ret := SortByKey(m).Desc().Values()

	t.Log(ret)
}

func TestJoin(t *testing.T) {
	m := map[string]any{
		"Bcd": "bcd",
		"Abc": "abc",
		"Efg": "efg",
	}

	j := SortByKey(m).Asc().Join("&", func(k string, v any) string {
		return fmt.Sprintf("%s=%v", k, v)
	})

	t.Log(j)
}
