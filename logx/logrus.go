package logx

import (
	"github.com/sirupsen/logrus"
)

type logrusWrpper struct {
	*logrus.Entry
}

func NewLogrus(opts ...logrusOption) *logrusWrpper {
	logger := logrus.New()

	logger.SetFormatter(NewNestedFormatter(NestedFormatterOption{}))

	for _, o := range opts {
		o(logger)
	}

	return &logrusWrpper{
		Entry: logrus.NewEntry(logger),
	}
}

func (l *logrusWrpper) SetLevel(level Level) {
	l.Logger.SetLevel(logrus.Level(level))
}

func (l *logrusWrpper) WithField(key string, value any) Logger {
	return l.WithFields(Fields{key: value})
}

func (l *logrusWrpper) WithFields(fields Fields) Logger {
	return &logrusWrpper{
		Entry: l.Logger.WithFields(logrus.Fields(fields)),
	}
}
