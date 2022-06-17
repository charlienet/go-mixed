package dateconv

import (
	"testing"
)

func TestParseDuration(t *testing.T) {
	t.Log(ParseDuration(""))
	t.Log(ParseDuration("abc"))
}
