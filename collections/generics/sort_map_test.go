package generics

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	m := map[string]any{
		"Bcd": "bcd",
		"Abc": "abc",
	}

	t.Log(NewSortMap(m).Values())
	t.Log(NewSortMap(m).Desc().Values())
	t.Log(NewSortMap(m).Asc().Values())
}

func TestMapSortInt(t *testing.T) {
	m := map[string]int{
		"Bcd": 8,
		"Abc": 4,
	}

	ret := NewSortMap(m).Desc().Values()
	t.Log(NewSortMap(m).Desc())
	t.Log(NewSortMap(m).Asc())

	t.Log(ret)
}

func TestJoin(t *testing.T) {
	m := map[string]any{
		"Bcd": "bcd",
		"Abc": "abc",
		"Efg": "efg",
	}
	t.Log(m)

	j := NewSortMap(m).Asc().Join("&", func(k string, v any) string {
		return fmt.Sprintf("%s=%v", k, v)
	})

	t.Log(j)
}
