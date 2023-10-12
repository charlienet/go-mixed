package compiledbuffer

import (
	"regexp"
	"strconv"
	"sync/atomic"
	"testing"
)

var s = "^aaa^"

func TestPutGet(t *testing.T) {
	b := NewCompiledBuffer(func(s string) (*regexp.Regexp, error) {
		t.Log("init")
		return regexp.Compile(s)
	})

	t.Log(b.Put(s))
	r, _ := b.Get(s)

	t.Log(r.Match([]byte("abc")))
	t.Log(r.Match([]byte("abc")))
}

func BenchmarkGet(b *testing.B) {
	b.Run("buf", func(b *testing.B) {
		buf := NewCompiledBuffer(func(s string) (*regexp.Regexp, error) { return regexp.Compile(s) })
		buf.Put(s)

		for i := 0; i < b.N; i++ {
			buf.Get(s)
		}
	})

	b.Run("buf-nocom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			regexp.Compile(s)
		}
	})
}

func BenchmarkConcurrence(b *testing.B) {
	buf := NewCompiledBuffer(func(s string) (*regexp.Regexp, error) { return regexp.Compile(s) })
	var i int64

	b.RunParallel(func(p *testing.PB) {
		gid := int(atomic.AddInt64(&i, 1) - 1)

		for i := 0; p.Next(); i++ {
			buf.Get(strconv.Itoa(gid))
		}
	})
}
