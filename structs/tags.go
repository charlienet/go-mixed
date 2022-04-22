package structs

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/charlienet/go-mixed/expr"
)

type tagOptions string

func parseTag(tag string) (string, tagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, tagOptions(opt)
}

func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}

type fieldOption struct {
	name      string
	omitEmpty bool
}

func getFieldOption(fi reflect.StructField) fieldOption {
	name, opts := parseTag(fi.Tag.Get(tagName))

	return fieldOption{
		name:      expr.If(isValidTag(name), name, fi.Name),
		omitEmpty: opts.Contains("omitempty"),
	}
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", c):
		case !unicode.IsLetter(c) && !unicode.IsDigit(c):
			return false
		}
	}
	return true
}
