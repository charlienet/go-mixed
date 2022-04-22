package logx_test

import (
	"testing"

	"github.com/charlienet/go-mixed/logx"
)

func TestBuilder(t *testing.T) {
	logger := logx.NewBuilder().
		WithLogrus().
		WithLogger()

	_ = logger
}
