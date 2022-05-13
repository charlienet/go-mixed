package structs_test

import (
	"reflect"
	"testing"

	"github.com/charlienet/go-mixed/structs"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	o := struct {
		Field1Name string
	}{Field1Name: "field 1 name"}

	s := structs.New(o)
	assert.Equal(t, reflect.Struct, s.Kind())

	t.Log(s.Names())
	t.Log(s.Values())
}

func TestIsZero(t *testing.T) {
	var v1 int
	assert.True(t, structs.IsZero(v1))

	var v2 = struct {
		Msg string
	}{}
	assert.True(t, structs.IsZero(v2))

	var v3 = struct {
		VV  int
		Msg string
	}{Msg: "abc"}
	assert.False(t, structs.IsZero(v3))

	v3.Msg = ""
	assert.True(t, structs.IsZero(v3))
}
