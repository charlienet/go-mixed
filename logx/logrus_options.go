package logx

import (
	"io"
	"log"
	"os"

	"github.com/charlienet/go-mixed/fs"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logrusOption func(logrusLogger *logrus.Logger)

type LogrusOptions struct {
	Level      string // 日志等级
	ShowCaller bool   // 是否记录调用者
}

type OutputType int

const (
	Console OutputType = iota
	File    OutputType = iota
	Both    OutputType = iota
)

type LogrusOutputOptions struct {
	FileName string              // 存储文件名
	Output   OutputType          // 输出选项
	Backup   LogrusBackupOptions // 文件备份选项
}

type LogrusBackupOptions struct {
	MaxSize    int  // 默认大小100M
	MaxAge     int  // 备份保留天数
	MaxBackups int  // 备份保留数量
	LocalTime  bool // 使用本地时间
	Compress   bool // 是否压缩备份
}

func (o LogrusBackupOptions) hasBackup() bool {
	return o.MaxSize > 0 ||
		o.MaxAge > 0 ||
		o.MaxBackups > 0 ||
		o.LocalTime ||
		o.Compress
}

func WithOptions(o LogrusOptions) logrusOption {
	return func(l *logrus.Logger) {
		level, err := logrus.ParseLevel(o.Level)
		if err != nil {
			level = logrus.TraceLevel
		}

		l.SetLevel(level)
		l.SetReportCaller(o.ShowCaller)
	}
}

func WithFormatter(formatter logrus.Formatter) logrusOption {
	return func(logrusLogger *logrus.Logger) {
		logrusLogger.SetFormatter(formatter)
	}
}

func WithOutput(options LogrusOutputOptions) logrusOption {
	return func(l *logrus.Logger) {
		var writer io.Writer
		switch {
		case options.Output == File && len(options.FileName) > 0:
			// 设置输出为文件，并且已经设置文件名
			writer = createFileWriter(options)
		case options.Output == Both && len(options.FileName) > 0:
			writer = io.MultiWriter(os.Stdout, createFileWriter(options))
		default:
			writer = os.Stdout
		}

		l.SetOutput(writer)
	}
}

func createFileWriter(options LogrusOutputOptions) io.Writer {
	if options.Backup.hasBackup() {
		return &lumberjack.Logger{
			Filename:   options.FileName,          // 日志文件名
			MaxSize:    options.Backup.MaxSize,    // 单位兆
			MaxBackups: options.Backup.MaxBackups, // 最多备份数量
			MaxAge:     options.Backup.MaxAge,     // 保留天数
			LocalTime:  options.Backup.LocalTime,  // 使用本地时间
			Compress:   options.Backup.Compress,   // 是否压缩
		}
	}

	f, err := fs.OpenOrNew(options.FileName)
	if err != nil {
		log.Panic(err)
	}

	return f
}
