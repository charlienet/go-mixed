package iprange

import "strings"

var maskBits = map[string]int{
	"255": 8,
	"254": 7,
	"252": 6,
	"248": 5,
	"240": 4,
	"224": 3,
	"192": 2,
	"128": 1,
	"0":   0,
}

func MaskToBits(mask string) int {
	bits := 0

	secs := strings.Split(mask, ".")
	if len(secs) != 4 {
		panic("the mask is incorrect")
	}

	for _, s := range secs {
		if v, ok := maskBits[s]; ok {
			bits += v
		} else {
			panic("the mask is incorrect")
		}
	}

	return bits
}
