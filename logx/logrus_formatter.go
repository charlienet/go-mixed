package logx

import (
	"fmt"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

const defaultTimestampFormat = "2006-01-02 15:04:05.000"

type NestedFormatterOption struct {
	Color bool
}

func NewJsonFormatter() logrus.Formatter {
	return &logrus.JSONFormatter{}
}

func NewTextFOrmatter() logrus.Formatter {
	return &logrus.TextFormatter{}
}

func NewNestedFormatter(option NestedFormatterOption) *nested.Formatter {
	return &nested.Formatter{
		TimestampFormat:       defaultTimestampFormat,
		NoColors:              !option.Color,
		CustomCallerFormatter: nestedCallerFormatter,
	}
}

func nestedCallerFormatter(f *runtime.Frame) string {
	_, filename := path.Split(f.File)
	return fmt.Sprintf(" (%s() %s:%d)", f.Function, filename, f.Line)
}
