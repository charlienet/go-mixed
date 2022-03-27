package crypto_test

import (
	"crypto/aes"
	"crypto/des"
	"encoding/hex"
	"testing"

	"github.com/charlienet/go-mixed/crypto"
)

func TestDes(t *testing.T) {
	key, _ := hex.DecodeString("0123456789ABCDEF")
	msg, _ := hex.DecodeString("F0A2B07E64DD2C25")

	c, err := crypto.Des(key).ECB().Encrypt(msg)
	t.Log(hex.EncodeToString(c), err)

	c, err = crypto.Des(key).Cbc().Encrypt(msg)
	t.Log(hex.EncodeToString(c), err)

	d, err := crypto.Des(key).Cbc().Decrypt(c)
	t.Log(hex.EncodeToString(d), err)
}

func TestTripleDES(t *testing.T) {
	key, _ := hex.DecodeString("0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF")
	block, err := des.NewTripleDESCipher(key)
	t.Log(block, err, block.BlockSize())
}

func TestAes(t *testing.T) {
	t.Log(aes.BlockSize)
	key, _ := hex.DecodeString("0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF")
	msg, _ := hex.DecodeString("F0A2B07E64DD2C25")
	c, err := crypto.Aes(key).ECB().Encrypt(msg)
	t.Log(hex.EncodeToString(c), err)
}
