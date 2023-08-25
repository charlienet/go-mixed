package mathx

import (
	"github.com/charlienet/go-mixed/expr"
	"golang.org/x/exp/constraints"
)

type Real interface {
	constraints.Integer | constraints.Float
}

// Max returns the larger one of v1 and v2.
func Max[T Real](v1, v2 T) T {
	return expr.Ternary(v1 > v2, v1, v2)
}

// Min returns the smaller one of v1 and v2.
func Min[T Real](v1, v2 T) T {
	return expr.Ternary(v1 < v2, v1, v2)
}

func Abs[T Real](val T) T {
	return expr.Ternary(val < 0, -val, val)
}

// Neg returns the negative of value. It does not negate value. For
// negating, simply use -value instead.
func Neg[T Real](value T) T {
	return expr.Ternary(value < 0, value, -value)
}

func Clamp[T Real](value, min, max T) T {
	if min > max {
		return min
	}

	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}
