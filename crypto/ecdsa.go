package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"
	"strconv"
)

// var _ Signer = &ecdsaOptions{}

type ecdsaOptions struct {
	hashOptions
	prv *ecdsa.PrivateKey
	pub *ecdsa.PublicKey
}

type Option interface {
	apply(*ecdsaOptions) error
}

type hash2Option interface {
	apply(*hashOptions) error
}

type privateKeyOption []byte

func (p privateKeyOption) apply(opts *ecdsaOptions) error {
	block, _ := pem.Decode(p)
	prv, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	opts.prv = prv

	return nil
}

type publicKeyOption []byte

func (p publicKeyOption) apply(opts *ecdsaOptions) error {
	block, _ := pem.Decode(p)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	opts.pub = pub.(*ecdsa.PublicKey)
	return nil
}

type hashOption crypto.Hash

func (p hashOption) apply(opts *hashOptions) error {
	opts.h = crypto.Hash(p)
	return nil
}

func NewEcdsa(h Hash, opts ...Option) (*ecdsaOptions, error) {
	i := &ecdsaOptions{}

	sh := crypto.Hash(h)
	if !sh.Available() {
		return nil, errors.New("unknown hash value " + strconv.Itoa(int(h)))
	}

	i.h = sh

	for _, v := range opts {
		if err := v.apply(i); err != nil {
			return nil, err
		}
	}

	return i, nil
}

func WithHash2(h Hash) hash2Option {
	return hashOption(h)
}

func ParsePrivateKey(pem []byte) Option {
	return privateKeyOption(pem)
}

func ParsePublicKey(pem []byte) Option {
	return publicKeyOption(pem)
}

func (opt *ecdsaOptions) Sign(msg []byte) ([]byte, error) {
	sum := opt.getHash(msg)
	return ecdsa.SignASN1(rand.Reader, opt.prv, sum)
}

func (opt *ecdsaOptions) Verify(msg, rText, sText []byte) bool {
	var r, s big.Int
	_ = r.UnmarshalText(rText)
	_ = s.UnmarshalText(sText)

	sum := opt.getHash(msg)
	return ecdsa.Verify(opt.pub, sum, &r, &s)
}
