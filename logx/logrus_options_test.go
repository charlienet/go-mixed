package logx_test

import (
	"os"
	"testing"

	"github.com/charlienet/go-mixed/logx"
)

func TestNewLogrus(t *testing.T) {
	_ = logx.WithOutput(logx.LogrusOutputOptions{
		Output: logx.Console,
	})
}

func TestFileOutput(t *testing.T) {
	t.Log(os.Getwd())
	logger := logx.NewLogrus(
		// logx.WithFormatter(&logrus.TextFormatter{}),
		logx.WithOptions(logx.LogrusOptions{ShowCaller: true}),
		logx.WithOutput(logx.LogrusOutputOptions{
			FileName: "abc.log",
		}))

	logger.Info("abc")
}
