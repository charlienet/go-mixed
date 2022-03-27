package crypto_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/charlienet/go-mixed/crypto"
)

func TestEsda(t *testing.T) {

	prv, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	t.Log(err)

	ecd, err := x509.MarshalECPrivateKey(prv)
	t.Log(err)

	secp256r1, _ := asn1.Marshal(asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7})
	fmt.Println(string(pem.EncodeToMemory(&pem.Block{Type: "EC PARAMETERS", Bytes: secp256r1})))
	b := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: ecd})
	fmt.Println(string(b))
}

func TestSign(t *testing.T) {
	ecdsa, err := crypto.NewEcdsa(crypto.SHA1)
	t.Log(err)

	_ = ecdsa
}
