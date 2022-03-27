package hash_test

import (
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/charlienet/go-mixed/hash"
)

func TestEncode(t *testing.T) {

	t.Log(base64.StdEncoding.EncodeToString(hash.Sha1([]byte{0x31})))
	t.Log(hex.EncodeToString(hash.Sha1([]byte{0x31})))

}

func TestXXHash(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(hex.EncodeToString(hash.XXhash([]byte(strconv.Itoa(i)))), "  ", hash.XXHashUint64([]byte(strconv.Itoa(i))))
	}
}

func TestMurmur3(t *testing.T) {
	t.Log(hash.Murmur3([]byte("123")))
	t.Log(hash.XXHashUint64([]byte("123")))
}
