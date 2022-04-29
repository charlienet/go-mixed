package structs

import (
	"reflect"
	"strings"

	"github.com/charlienet/go-mixed/expr"
)

type field struct {
	name        string
	tagName     string
	ignoreEmpty bool
	ignore      bool
}

func parseField(fi reflect.StructField, opt option) field {
	name, opts := parseTag(fi.Tag.Get(opt.TagName))

	return field{
		name:        fi.Name,
		tagName:     expr.If(isValidTag(name), name, expr.If(opt.NameConverter != nil, opt.NameConverter(fi.Name), fi.Name)),
		ignoreEmpty: opt.IgnoreEmpty || (opts.Contains("omitempty") && opt.Omitempty),
		ignore:      (name == "-" && opt.Ignore) || isSkipField(fi.Name, opt.SkipFields),
	}
}

func (f field) shouldIgnore(s reflect.Value) bool {
	return f.ignore || (s.IsZero() && f.ignoreEmpty)
}

func isSkipField(name string, skips []string) bool {
	for _, v := range skips {
		if strings.EqualFold(v, name) {
			return true
		}
	}

	return false
}
