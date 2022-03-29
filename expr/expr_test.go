package expr

import "testing"

func TestIf(t *testing.T) {
	v1 := 10
	v2 := 4
	t.Log(If(v1 > v2, v1, v2))
}
