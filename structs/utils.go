package structs

import (
	"reflect"

	"github.com/charlienet/go-mixed/json"
)

type optionFunc func(*option)

type option struct {
	NameFunc    func(string) string
	IgnoreEmpty bool
	DeepCopy    bool
	Omitempty   bool
}

func IgnoreEmpty() optionFunc {
	return func(o *option) {
		o.IgnoreEmpty = true
	}
}

func DeepCopy() optionFunc {
	return func(o *option) {
		o.DeepCopy = true
	}
}

func Omitempty() optionFunc {
	return func(o *option) {
		o.Omitempty = true
	}
}

func Lcfirst() optionFunc {
	return func(o *option) {
		o.NameFunc = json.Lcfirst
	}
}

func Camel2Case() optionFunc {
	return func(o *option) {
		o.NameFunc = json.Camel2Case
	}
}

func createOptions(opts []optionFunc) *option {
	o := &option{
		NameFunc: func(s string) string { return s },
	}

	for _, f := range opts {
		f(o)
	}

	return o
}

func shouldIgnore(v reflect.Value, ignoreEmpty bool) bool {
	return ignoreEmpty && v.IsZero()
}
