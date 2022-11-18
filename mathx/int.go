package mathx

import (
	"github.com/charlienet/go-mixed/expr"
	"golang.org/x/exp/constraints"
	"unsafe"
)

// Max returns the larger one of v1 and v2.
func Max[T constraints.Ordered](v1, v2 T) T {
	return expr.Ternary(v1 > v2, v1, v2)
}

// Min returns the smaller one of v1 and v2.
func Min[T constraints.Ordered](v1, v2 T) T {
	return expr.Ternary(v1 < v2, v1, v2)
}

func Abs1[T constraints.Signed](n T) T {
	shift := 63
	switch unsafe.Sizeof(n) {
	case 1:
		shift = 7
	case 4:
		shift = 31
	}

	y := n >> shift
	return T((n ^ y) - y)
}

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
