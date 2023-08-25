package stringx_test

import (
	"testing"

	"github.com/charlienet/go-mixed/stringx"
	"github.com/stretchr/testify/assert"
)

func TestMask(t *testing.T) {
	cases := []struct {
		origin   string
		place    stringx.Place
		length   int
		excepted string
	}{
		{"aa", stringx.Begin, 6, "**"},
		{"18980832408", stringx.Begin, 4, "****0832408"},
		{"18980832408", stringx.End, 4, "1898083****"},
		{"18980832408", stringx.Middle, 4, "189****2408"},
	}

	a := assert.New(t)
	for _, c := range cases {
		t.Log(stringx.Mask(c.origin, c.place, c.length))
		a.Equal(c.excepted, stringx.Mask(c.origin, c.place, c.length))
	}
}
