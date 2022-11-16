package mathx

import (
	"github.com/charlienet/go-mixed/expr"
	"golang.org/x/exp/constraints"
)

// MaxInt returns the larger one of v1 and v2.
func Max[T constraints.Ordered](v1, v2 T) T {
	return expr.Ternary(v1 > v2, v1, v2)
}

// MinInt returns the smaller one of v1 and v2.
func Min[T constraints.Ordered](v1, v2 T) T {
	return expr.Ternary(v1 < v2, v1, v2)
}

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
