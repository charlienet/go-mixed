package fs

import (
	"io"
	"os"
	"path/filepath"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}

	return file.IsDir()
}

// 打开或新建文件，目录不存在时创建目录
func OpenOrNew(filename string) (io.Writer, error) {
	dir := filepath.Dir(filename)
	if !IsExist(dir) {
		os.MkdirAll(dir, 0744)
	}

	mode := os.FileMode(0644)
	return os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, mode)
}
