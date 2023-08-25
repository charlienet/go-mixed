package mathx_test

import (
	"testing"

	"github.com/charlienet/go-mixed/mathx"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, mathx.Min(1, 3))
	assert.Equal(t, 2, mathx.Min(66, 2))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 3, mathx.Max(1, 3))
	assert.Equal(t, 66, mathx.Max(66, 2))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 23, mathx.Abs(23))
	assert.Equal(t, 23, mathx.Abs(-23))
	assert.Equal(t, 0, mathx.Abs(0))

	var u int8 = -127
	var exp int8 = 127
	assert.Equal(t, exp, mathx.Abs(u))

	assert.Equal(t, 1.23, mathx.Abs(-1.23))
}

func TestClamp(t *testing.T) {
	tests := []struct{ value, min, max, want int }{
		// Min.
		{-1, 0, 2, 0},
		{0, 0, 2, 0},
		// Mid.
		{1, 0, 2, 1},
		{2, 0, 2, 2},
		// Max.
		{2, 0, 2, 2},
		{3, 0, 2, 2},
		// Empty range.
		{-1, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
	}
	for _, test := range tests {
		got := mathx.Clamp(test.value, test.min, test.max)
		if got != test.want {
			t.Errorf("Clamp(%v, %v, %v) = %v, want %v", test.value, test.min, test.max, got, test.want)
		}
	}
}
