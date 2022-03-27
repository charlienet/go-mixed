package crypto_test

import (
	"encoding/hex"
	"testing"

	"github.com/charlienet/go-mixed/crypto"
	"github.com/tjfoc/gmsm/sm4"
)

func TestGmsmSm4(t *testing.T) {
	key, _ := hex.DecodeString("0123456789ABCDEFFEDCBA9876543210")
	msg, _ := hex.DecodeString("F0A2B07E64DD2C2590F93E4EDD90FBB4")

	c, err := sm4.Sm4Ecb(key, msg, true)
	t.Log(hex.EncodeToString(c), err)

	d, err := sm4.Sm4Ecb(key, c, false)
	t.Log(hex.EncodeToString(d), err)
}

func TestPadding(t *testing.T) {
	msg, _ := hex.DecodeString("F0A2B07E64DD2C2590F93E4EDD90FBB4")

	blockSize := sm4.BlockSize
	padding := blockSize - len(msg)%blockSize
	t.Log(padding)
}

func TestSm4(t *testing.T) {
	key := []byte("1234567890abcdef")
	msg := []byte("123321123321123321123321123321")

	cipherText, err := crypto.Sm4(key).ECB().Encrypt(msg)
	t.Log("ECB加密:", hex.EncodeToString(cipherText), err)

	de, err := crypto.Sm4(key).ECB().Decrypt(cipherText)
	t.Log("ECB解密:", string(de), err)

	cipherText, err = crypto.Sm4(key).CBC().WithIV(key).Encrypt(msg)
	t.Log("CBC加密:", hex.EncodeToString(cipherText), err)

	de, err = crypto.Sm4(key).CBC().WithIV(key).Decrypt(cipherText)
	t.Log("CBC解密:", string(de), err)
}

func TestDecrypt(t *testing.T) {
	key := []byte("XbBpuLSzaXtlOYFV")
	iv := []byte("UISwD9fW6cFh9SNS")
	en := "BAD4C05DB0A51895A38D976F97057C2D1743473CE6DABC3456DD4EA751A9794D81096050DBA084F1CB3791C63DFFEDD1D63B046B155FD06386DEE8434A20D8A7465780EF3660ED1073A253DEA4768AB735E2DDEB4602927D3FF85E429C9B7557E6A3A198F4781642CDD30449968FBD2E54E0425E327805DFB0A1DA4FAE33AC68A3377D20042A9459EEF09BEE8CBE483BF61D32B7BB402730AA2276EA3C3A078B895D684A91DD7EEF0F7A25289B1D4905AF524126E8C3DBCB0AB73C92ABC1A83ECA687777B9B609DD8B0F69602EC3E74243E00B33D51EDF930A5316BCB388E4B7B2A6EFDD8B0BE4A19625D297B25D2BD2E5424F2E9B6A4BBF6A70DBE3C6ABB635554AC21CE053D7ECA23D82EF8060C874D507FC27CFCC06EDF41AF98ED0C2C59E39146CC28BA7630D74870BD372863FC4"

	c := crypto.Sm4(key).CBC().WithIV(iv)

	b, _ := hex.DecodeString(en)
	decrypted, err := c.Decrypt(b)
	t.Log(hex.EncodeToString(sm4.IV))
	t.Log(hex.EncodeToString(decrypted))
	t.Log(string(decrypted), err)

	encrypted, err := c.Encrypt(decrypted)
	t.Log(err)
	ddd, err := c.Decrypt(encrypted)
	t.Log(string(ddd), err)
}
