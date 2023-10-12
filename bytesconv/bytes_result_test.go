package bytesconv_test

import (
	"testing"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/rand"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
)

func TestHexUppercase(t *testing.T) {
	b, _ := rand.RandBytes(12)

	l := bytesconv.BytesResult(b).Hex()
	t.Log(l)

	u := bytesconv.BytesResult(b).UppercaseHex()
	t.Log(u)
}

func TestHexToBase64(t *testing.T) {
	v := "abc"

	unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)

	t.Log(bytesconv.FromString(v).Base64())

	b, _ := rand.RandBytes(43)
	t.Log(bytesconv.FromBytes(b).Base64())

	c, n := charset.Lookup("utf8")

	
	t.Log(c, n)

}
