package stringx

import (
	"strings"
)

type Place int

const (
	Begin Place = iota
	Middle
	End
)

func Mask(s string, place Place, length int, mask ...rune) string {
	m := '*'
	if len(mask) > 0 {
		m = mask[0]
	}

	n := len(s)
	if length >= n {
		return strings.Repeat(string(m), n)
	}

	i := 0
	if place == Middle {
		i = (n - length) / 2
	} else if place == End {
		i = n - length
	}

	end := i + length
	r := []rune(s)
	for ; i < end; i++ {
		r[i] = m
	}

	return string(r)
}
