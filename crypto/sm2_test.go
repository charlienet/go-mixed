package crypto_test

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/charlienet/go-mixed/crypto"
)

func TestNewSm2(t *testing.T) {
	o, err := crypto.NewSm2()
	t.Logf("%+v, %v", o, err)

	t.Log(crypto.NewSm2(crypto.ParseSm2PrivateKey([]byte{}, []byte{})))

	msg := []byte("123456")
	sign, err := o.Sign(msg)
	t.Log(hex.EncodeToString(sign), err)

	ok := o.Verify(msg, sign)
	if !ok {
		t.Fail()
	}
	t.Log(ok)
}

const (
	privPem = `-----BEGIN ENCRYPTED PRIVATE KEY-----
MIH8MFcGCSqGSIb3DQEFDTBKMCkGCSqGSIb3DQEFDDAcBAgXsd3MYu0BwwICCAAw
DAYIKoZIhvcNAgcFADAdBglghkgBZQMEASoEEJzb8/1Aqhbv2cf777VoW0cEgaAz
DbRJgs76YYpya9wiaZeAavSn8Ydi+CYSvvQurqa1q0Hmna/Lgcgt2Z0F3fFN/EYP
wmDCd6SQ5hdPfQLBtkpDQdFylIHAm26O0smciB7NlfWSdgIluFacbMJ++/YHvcDp
yl1qcRpjk+s+1+8YBUp7Mp1CXbDXdQebH9xezOE3OH8+9zO3qi5qeLEVofgRQJIY
k8EBbLsGMy4WlSr0u29A
-----END ENCRYPTED PRIVATE KEY-----`

	pubPem = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEvfHGxZL/wzWLYgPsHEpFxCCwXKSr
XExvTJS6FAem+lQTyHwOGT+qFf67J77d5y/exn6E5br79nsJkoM/7A72nQ==
-----END PUBLIC KEY-----`

	badPubPem = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE3Og1rzeSs2wO9+YFIdgnAES03u1n
hslcifiQY8173nHtaB3R6T0PwRQTwKbpdec0dwVCpvVcdzHtivndlG0mqQ==
-----END PUBLIC KEY-----`
)

func TestPrivatePem(t *testing.T) {
	signer, err := crypto.NewSm2(
		crypto.ParseSm2PrivateKey([]byte(privPem), []byte{}),
		crypto.ParseSm2PublicKey([]byte(pubPem)))

	t.Log(signer, err)
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}

	msg := []byte("123456")
	sign, err := signer.Sign(msg)
	t.Log(hex.EncodeToString(sign), err)

	t.Log(signer.Verify(msg, sign))
}

func TestBadPublicPem(t *testing.T) {
	signer, err := crypto.NewSm2(
		crypto.ParseSm2PrivateKey([]byte(privPem), []byte{}),
		crypto.ParseSm2PublicKey([]byte(badPubPem)))

	t.Log(signer, err)

	msg := []byte("123456")
	sign, err := signer.Sign(msg)
	t.Log(hex.EncodeToString(sign), err)

	t.Log(signer.Verify(msg, sign))
}

const pemString = `-----BEGIN EC PARAMETERS-----
BggqgRzPVQGCLQ==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAU/RPiFOw8sI+4dM/0ZusJ7dWxi72DpnOukgGNZfPP5oAoGCCqBHM9V
AYItoUQDQgAEbl5hPO00SJnkTpNjefes6QjmOrhQTrcocBQ0V9yB3ow/COroyHIp
MV8UROLaT5kNUim8Z6XQjL+TWrfo11JQ2w==
-----END EC PRIVATE KEY-----`

func TestDecodePem(t *testing.T) {

	block, _ := pem.Decode([]byte(pemString))
	fmt.Println(string(block.Bytes))

	prv, err := x509.ParseECPrivateKey(block.Bytes)
	t.Log(prv, err)
}
