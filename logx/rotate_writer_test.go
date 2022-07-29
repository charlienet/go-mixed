package logx

import (
	"path/filepath"
	"testing"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func TestNewWriter(t *testing.T) {
	t.Log(filepath.Abs("logs"))

	logf, err := rotatelogs.New("logs/aaaa.%Y%m%d.log",
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))

	if err != nil {
		t.Fatal(err)
	}
	defer logf.Close()

	t.Log(logf.CurrentFileName())

	_, err = logf.Write([]byte("abaccad"))
	t.Log(err)
}
