package hmac_test

import (
	"crypto/hmac"
	"encoding/hex"
	"testing"

	"github.com/tjfoc/gmsm/sm3"
)

func TestHmac(t *testing.T) {

	key := []byte("")
	msg := []byte("123")

	h := hmac.New(sm3.New, key)
	h.Write(msg)
	t.Log(hex.EncodeToString(h.Sum(nil)))
}
