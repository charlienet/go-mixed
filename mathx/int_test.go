package mathx_test

import (
	"github.com/charlienet/go-mixed/mathx"
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.Equal(t, 23, mathx.Abs1(23))
	assert.Equal(t, 23, mathx.Abs1(-23))
	assert.Equal(t, 0, mathx.Abs1(0))

	var u int8 = -127
	var exp int8 = 127
	assert.Equal(t, exp, mathx.Abs1(u))
}
