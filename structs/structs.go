package structs

import (
	"errors"
	"reflect"
)

const defaultTagName = "json"

var (
	ErrInvalidCopyDestination = errors.New("copy destination is invalid")
)

type Struct struct {
	opt    option
	raw    any
	value  reflect.Value
	fields []field
}

func New(o any, opts ...optionFunc) *Struct {
	opt := acquireOptions(opts)
	v := indirect(reflect.ValueOf(o))

	return &Struct{
		opt:    opt,
		raw:    o,
		value:  v,
		fields: parseFields(v, opt),
	}
}

func (s *Struct) Kind() reflect.Kind {
	return s.value.Kind()
}

func (s *Struct) Names() []string {
	names := make([]string, len(s.fields))
	for i, f := range s.fields {
		names[i] = f.name
	}

	return names
}

func (s *Struct) Values() []any {
	values := make([]any, 0, len(s.fields))
	for _, fi := range s.fields {
		v := s.value.FieldByName(fi.name)
		values = append(values, v.Interface())
	}

	return values
}

func (s *Struct) ToMap() map[string]any {
	m := make(map[string]any, len(s.fields))
	for _, fi := range s.fields {
		source := s.value.FieldByName(fi.name)
		if fi.shouldIgnore(source) {
			continue
		}

		m[fi.tagName] = source.Interface()
	}

	return m
}

func (s *Struct) Copy(dest any) error {
	to := indirect(reflect.ValueOf(dest))

	if !to.CanAddr() {
		return ErrInvalidCopyDestination
	}

	t := indirectType(reflect.TypeOf(dest))
	for i := 0; i < t.NumField(); i++ {
		destField := t.Field(i)
		if fi, ok := s.getByName(destField.Name); ok {
			source := s.value.FieldByName(fi.name)
			if fi.shouldIgnore(source) {
				continue
			}

			tv := to.FieldByName(destField.Name)
			tv.Set(source)
		}
	}

	return nil
}

func (s *Struct) getByName(name string) (field, bool) {
	for i := range s.fields {
		f := s.fields[i]
		if f.name == name {
			return f, true
		}
	}

	return field{}, false
}

func ToMap(o any, opts ...optionFunc) map[string]any {
	return New(o, opts...).ToMap()
}

func Copy(source, dst any, opts ...optionFunc) {
	New(source, opts...).Copy(dst)
}

func parseFields(t reflect.Value, opt option) []field {
	typ := indirectType(t.Type())
	num := typ.NumField()
	fields := make([]field, 0, num)
	for i := 0; i < num; i++ {
		fi := typ.Field(i)
		fields = append(fields, parseField(fi, opt))
	}

	return fields
}

func indirect(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}

func indirectType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		return t.Elem()
	}

	return t
}
