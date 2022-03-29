package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"

	"github.com/tjfoc/gmsm/sm3"
)

func Md5(key, msg []byte) hashResult { return sum(md5.New, key, msg) }

func Sha1(key, msg []byte) hashResult { return sum(sha1.New, key, msg) }

func Sha224(key, msg []byte) hashResult { return sum(sha256.New224, key, msg) }

func Sha256(key, msg []byte) hashResult { return sum(sha256.New, key, msg) }

func Sha384(key, msg []byte) hashResult { return sum(sha512.New384, key, msg) }

func Sha512(key, msg []byte) hashResult { return sum(sha512.New, key, msg) }

func Sm3(key, msg []byte) hashResult { return sum(sm3.New, key, msg) }

func sum(f func() hash.Hash, msg, key []byte) hashResult {
	h := hmac.New(f, key)

	h.Write(msg)
	return h.Sum(nil)
}

type hashResult []byte

func (r hashResult) Hex() string {
	return hex.EncodeToString(r)
}

func (r hashResult) Base64() string {
	return base64.StdEncoding.EncodeToString(r)
}

func (r hashResult) Bytes() []byte {
	return r
}

func (r hashResult) String() string {
	return r.Hex()
}
