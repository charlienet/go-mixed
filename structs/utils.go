package structs

import (
	"reflect"
)

type optionFunc func(*option)

type option struct {
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

func createOptions(opts []optionFunc) *option {
	o := &option{}
	for _, f := range opts {
		f(o)
	}

	return o
}

func shouldIgnore(v reflect.Value, ignoreEmpty bool) bool {
	return ignoreEmpty && v.IsZero()
}

