package json

import (
	"testing"
)

func TestMapModify(t *testing.T) {
	m := map[string]string{
		"A": "A",
	}

	modify(m)
	t.Log(m)
}

func modify(m map[string]string) {
	m["B"] = "bbb"
}
