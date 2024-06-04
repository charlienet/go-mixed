package redis

import (
	"strings"

	"github.com/charlienet/go-mixed/expr"
)

const (
	defaultSeparator = ":"
)

type redisPrefix struct {
	prefix    string
	separator string
}

func newPrefix(separator string, prefix ...string) redisPrefix {
	s := expr.Ternary(len(separator) == 0, defaultSeparator, separator)

	return redisPrefix{
		separator: s,
		prefix:    expr.Ternary(len(prefix) > 0, strings.Join(prefix, separator), ""),
	}
}

func (p *redisPrefix) Prefix() string {
	return p.prefix
}

func (p *redisPrefix) Separator() string {
	return p.separator
}

func (p *redisPrefix) hasPrefix() bool {
	return len(p.prefix) > 0
}

func (p *redisPrefix) join(key ...string) string {
	s := make([]string, 0, len(key)+1)
	s = append(s, p.prefix)
	s = append(s, key...)

	return strings.Join(s, p.separator)
}
