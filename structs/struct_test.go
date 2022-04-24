package structs_test

import (
	"reflect"
	"testing"

	"github.com/charlienet/go-mixed/structs"
	"github.com/go-playground/assert/v2"
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
