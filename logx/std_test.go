package logx_test

import (
	"testing"

	"github.com/charlienet/go-mixed/logx"
)

func TestStandardLogger(t *testing.T) {
	l := logx.StandardLogger()

	l.Debug("abc")
	l.Info("abc")
	l.Warn("abc")
	l.Error("abc")
}
