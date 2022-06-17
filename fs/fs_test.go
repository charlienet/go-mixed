package fs

import "testing"

func TestOpenFile(t *testing.T) {
	OpenOrNew("logs/aaa.log")
}
