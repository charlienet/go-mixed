package logx

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogrusInfo(t *testing.T) {
	l := NewLogrus()
	l.Info("abcdef")
}

func TestWithField(t *testing.T) {
	var l Logger = NewLogrus(WithFormatter(NewNestedFormatter(NestedFormatterOption{Color: false})))

	l.WithField("aaa", "aaa").Info("aaaa")
	l.WithField("bbb", "bbb").Info("bbbb")

	l = l.WithField("111", "111")
	l.WithField("222", "222").Info("222")
}

func TestLogrus(t *testing.T) {
	l := logrus.NewEntry(logrus.New())
	l.WithField("abc", "bcd").Info("aaaa")
	l.WithField("bbb", "bbb").Info("bbbb")
	l.WithField("ccc", "ccc").Info("cccc")
}

func TestLevel(t *testing.T) {
	logger := NewLogrus()
	logger.Info("abcd")

	// l, _ := ParseLevel("Warn")
	// logger.SetLevel(l)
	logger.Info("bcdefg")
}

func TestMutiWriter(t *testing.T) {
	l := NewLogger().AppendLogger()

	_ = l
}
