package structs

import (
	"reflect"
)

const tagName = "json"

func ToMap(o any, opts ...optionFunc) map[string]any {
	typ := reflect.TypeOf(o)

	kind := typ.Kind()
	if kind == reflect.Map {
		if h, ok := o.(map[string]any); ok {
			return h
		}
	}

	opt := createOptions(opts)
	val := reflect.ValueOf(o)
	m := make(map[string]any)
	for i := 0; i < val.NumField(); i++ {
		fi := typ.Field(i)

		field := getFieldOption(fi)
		source := val.FieldByName(fi.Name)
		if shouldIgnore(source, opt.IgnoreEmpty || field.omitEmpty && opt.Omitempty) {
			continue
		}

		m[opt.NameFunc(field.name)] = source.Interface()
	}

	return m
}
