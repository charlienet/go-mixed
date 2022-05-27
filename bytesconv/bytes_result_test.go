package bytesconv_test

import (
	"testing"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/rand"
)

func TestHexUppercase(t *testing.T) {
	b, _ := rand.RandBytes(12)

	l := bytesconv.BytesResult(b).Hex()
	t.Log(l)

	u := bytesconv.BytesResult(b).UppercaseHex()
	t.Log(u)
}
