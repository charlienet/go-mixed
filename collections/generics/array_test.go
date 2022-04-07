package generics

import "testing"

func TestArray(t *testing.T) {
	a := NewArray(9, 1, 2, 4, 3, 1, 3)
	t.Log(a)

	b := a.Distinct(true)
	t.Log(b)
	t.Log(b.ToList())
}
