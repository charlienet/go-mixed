package structs

import "testing"

func TestCopy(t *testing.T) {
	v1 := struct{ Abc string }{Abc: "abc"}
	v2 := struct{ Abc string }{}

	Copy(v1, &v2, IgnoreEmpty())

	t.Log(v2)
}
