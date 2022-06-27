package logx

import "io"

var std = defaultLogger()

func StandardLogger() Logger {
	return std
}

func defaultLogger() Logger {
	return NewLogrus(
		WithOptions(LogrusOptions{Level: "Debug"}),
		WithOutput(LogrusOutputOptions{}),
		WithFormatter(NewNestedFormatter(NestedFormatterOption{Color: true})))
}

// Fields type, used to pass to `WithFields`.
type Fields map[string]any

type Logger interface {
	SetLevel(level Level)
	WithField(key string, value any) Logger
	WithFields(fields Fields) Logger
	Trace(args ...any)
	Tracef(format string, args ...any)
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Println(args ...any)
	Print(args ...any)
	Printf(format string, args ...any)
	Writer() io.Writer
}
