package structs

import (
	"github.com/charlienet/go-mixed/json"
)

type optionFunc func(*option)

type option struct {
	SkipFields    []string
	TagName       string
	DeepCopy      bool
	Omitempty     bool
	IgnoreEmpty   bool
	Ignore        bool
	NameConverter func(string) string
}

func TagName(name string) optionFunc {
	return func(o *option) {
		o.TagName = name
	}
}

func IgnoreEmpty() optionFunc {
	return func(o *option) {
		o.IgnoreEmpty = true
	}
}

func Omitempty() optionFunc {
	return func(o *option) {
		o.Omitempty = true
	}
}

func DeepCopy() optionFunc {
	return func(o *option) {
		o.DeepCopy = true
	}
}

func SkipField(field string) optionFunc {
	return func(o *option) {
		o.SkipFields = append(o.SkipFields, field)
	}
}

func SkipFields(fields []string) optionFunc {
	return func(o *option) {
		o.SkipFields = append(o.SkipFields, fields...)
	}
}

func Lcfirst() optionFunc {
	return func(o *option) {
		o.NameConverter = json.Lcfirst
	}
}

func Camel2Case() optionFunc {
	return func(o *option) {
		o.NameConverter = json.Camel2Case
	}
}

func defaultOptions() option {
	return option{
		TagName:       defaultTagName,
		Ignore:        true,
		NameConverter: func(s string) string { return s },
	}
}

func acquireOptions(opts []optionFunc) option {
	o := defaultOptions()
	for _, f := range opts {
		f(&o)
	}

	return o
}
