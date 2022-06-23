package hmac

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"strings"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/crypto"
	"github.com/tjfoc/gmsm/sm3"
)

var _ crypto.Signer = &hashComparer{}

type HMacFunc func(key, msg []byte) bytesconv.BytesResult

var hmacFuncs = map[string]HMacFunc{
	"HmacMD5":    Md5,
	"HmacSHA1":   Sha1,
	"HmacSHA224": Sha224,
	"HmacSHA256": Sha256,
	"HmacSHA384": Sha384,
	"HmacSHA512": Sha512,
	"HmacSM3":    Sm3,
}

type hashComparer struct {
	key      []byte
	hashFunc HMacFunc
}

func New(fname string, key []byte) (*hashComparer, error) {
	f, err := ByName(fname)
	if err != nil {
		return nil, err
	}

	return &hashComparer{
		key:      key,
		hashFunc: f,
	}, nil
}

func (c *hashComparer) Sign(msg []byte) ([]byte, error) {
	ret := c.hashFunc(c.key, msg)
	return ret.Bytes(), nil
}

func (c *hashComparer) Verify(msg, target []byte) bool {
	ret := c.hashFunc(c.key, msg)

	return bytes.Compare(ret.Bytes(), target) == 0
}

func ByName(name string) (HMacFunc, error) {
	if f, ok := hmacFuncs[strings.ToUpper(name)]; ok {
		return f, nil
	}

	return nil, errors.New("Unsupported hash functions")
}

func Md5(key, msg []byte) bytesconv.BytesResult { return sum(md5.New, key, msg) }

func Sha1(key, msg []byte) bytesconv.BytesResult { return sum(sha1.New, key, msg) }

func Sha224(key, msg []byte) bytesconv.BytesResult { return sum(sha256.New224, key, msg) }

func Sha256(key, msg []byte) bytesconv.BytesResult { return sum(sha256.New, key, msg) }

func Sha384(key, msg []byte) bytesconv.BytesResult { return sum(sha512.New384, key, msg) }

func Sha512(key, msg []byte) bytesconv.BytesResult { return sum(sha512.New, key, msg) }

func Sm3(key, msg []byte) bytesconv.BytesResult { return sum(sm3.New, key, msg) }

func sum(f func() hash.Hash, msg, key []byte) bytesconv.BytesResult {
	h := hmac.New(f, key)

	h.Write(msg)
	return h.Sum(nil)
}
