package logx

import (
	"fmt"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type logrusWrpper struct {
	logrus *logrus.Entry
}

func NewLogrus() Logger {
	logger := logrus.New()

	logger.SetFormatter(
		&nested.Formatter{
			TimestampFormat:       "2006-01-02 15:04:05.999",
			NoColors:              false,
			CustomCallerFormatter: nestedCallerFormatter,
		})

	return &logrusWrpper{
		logrus: logrus.NewEntry(logger),
	}
}

func nestedCallerFormatter(f *runtime.Frame) string {
	_, filename := path.Split(f.File)
	return fmt.Sprintf(" (%s() %s:%d)", f.Function, filename, f.Line)
}

func (l *logrusWrpper) SetLevel(level Level) {
	l.logrus.Logger.SetLevel(logrus.Level(level))
}

func (l *logrusWrpper) WithField(key string, value any) Logger {
	return l.WithFields(Fields{key: value})
}

func (l *logrusWrpper) WithFields(fields Fields) Logger {
	return &logrusWrpper{
		logrus: l.logrus.WithFields(logrus.Fields(fields)),
	}
}

func (l *logrusWrpper) Trace(args ...any) {
	l.logrus.Trace(args...)
}

func (l *logrusWrpper) Tracef(format string, args ...any) {
	l.logrus.Tracef(format, args...)
}

func (l *logrusWrpper) Debug(args ...any) {
	l.logrus.Debug(args...)
}

func (l *logrusWrpper) Debugf(format string, args ...any) {
	l.logrus.Debugf(format, args...)
}

func (l *logrusWrpper) Info(args ...any) {
	l.logrus.Info(args...)
}

func (l *logrusWrpper) Infof(format string, args ...any) {
	l.logrus.Infof(format, args...)
}

func (l *logrusWrpper) Warn(args ...any) {
	l.logrus.Warn(args...)
}

func (l *logrusWrpper) Error(args ...any) {
	l.logrus.Error(args...)
}

func (l *logrusWrpper) Warnf(format string, args ...any) {
	l.logrus.Warnf(format, args...)
}

func (l *logrusWrpper) Errorf(format string, args ...any) {
	l.logrus.Errorf(format, args...)
}

func (l *logrusWrpper) Fatal(args ...any) {
	l.logrus.Fatal(args)
}

func (l *logrusWrpper) Fatalf(format string, args ...any) {
	l.logrus.Fatalf(format, args...)
}

func (l *logrusWrpper) Println(args ...any) {
	l.logrus.Infoln(args...)
}

func (l *logrusWrpper) Print(args ...any) {
	l.logrus.Info(args...)
}

func (l *logrusWrpper) Printf(format string, args ...any) {
	l.logrus.Infof(format, args...)
}
