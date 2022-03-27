package fs

import (
	"os"
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
