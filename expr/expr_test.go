package expr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	v1 := 10
	v2 := 4
	t.Log(Ternary(v1 > v2, v1, v2))
}

func TestIf(t *testing.T) {
	is := assert.New(t)

	is.Equal(1, If(true, 1).ElseIf(false, 2).Else(3))
	is.Equal(1, If(true, 1).ElseIf(true, 2).Else(3))
	is.Equal(2, If(false, 1).ElseIf(true, 2).Else(3))
	is.Equal(3, If(false, 1).ElseIf(false, 2).Else(3))
}

func TestSwitch(t *testing.T) {
	is := assert.New(t)

	is.Equal(1, Switch[int, int](42).Case(42, 1).Case(1, 2).Default(3))
	is.Equal(1, Switch[int, int](42).Case(42, 1).Case(42, 2).Default(3))
	is.Equal(1, Switch[int, int](42).Case(1, 1).Case(42, 2).Default(3))
	is.Equal(1, Switch[int, int](42).Case(1, 1).Case(1, 2).Default(3))
}
