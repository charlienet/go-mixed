package logx

import (
	"fmt"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

type logrusWrpper struct {
	logger *logrus.Entry
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
		logger: logrus.NewEntry(logger),
	}
}

func nestedCallerFormatter(f *runtime.Frame) string {
	_, filename := path.Split(f.File)
	return fmt.Sprintf(" (%s() %s:%d)", f.Function, filename, f.Line)
}

func (l *logrusWrpper) SetLevel(level Level) {
	l.logger.Logger.SetLevel(logrus.Level(level))
}

func (l *logrusWrpper) WithField(key string, value any) Logger {
	return l.WithFields(Fields{key: value})
}

func (l *logrusWrpper) WithFields(fields Fields) Logger {
	return &logrusWrpper{
		logger: l.logger.WithFields(logrus.Fields(fields)),
	}
}

func (l *logrusWrpper) Info(args ...any) {
	l.logger.Info(args...)
}

func (l *logrusWrpper) Infof(format string, args ...any) {
	l.logger.Infof(format, args...)
}

func (l *logrusWrpper) Warn(args ...any) {
	l.logger.Warn(args...)
}

func (l *logrusWrpper) Error(args ...any) {
	l.logger.Error(args...)
}

func (l *logrusWrpper) Warnf(format string, args ...any) {
	l.logger.Warnf(format, args...)
}

func (l *logrusWrpper) Errorf(format string, args ...any) {
	l.logger.Errorf(format, args...)
}

func (l *logrusWrpper) Fatalf(format string, args ...any) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusWrpper) Println(args ...any) {
	l.logger.Infoln(args...)
}

func (l *logrusWrpper) Print(args ...any) {
	l.logger.Info(args...)
}

func (l *logrusWrpper) Printf(format string, args ...any) {
	l.logger.Infof(format, args...)
}
