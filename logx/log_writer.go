package logx

import (
	"io"
	"os"
)

type Rotate int

const (
	None Rotate = iota // 不分割日志
	Size               // 按大小分割
	Date               // 按日期分割
)

type OutputOptions struct {
	LogrusOutputOptions
}

func WithFile(filename string) (io.Writer, error) {
	mode := os.FileMode(0644)
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, mode)
}
