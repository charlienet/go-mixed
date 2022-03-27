package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"github.com/tjfoc/gmsm/sm3"
)

func Sha1(key, msg []byte) []byte { return sum(sha1.New, key, msg) }

func Sha1Hex(key, msg []byte) string { return hex.EncodeToString(Sha1(key, msg)) }

func Sha1Base64(key, msg []byte) string { return hex.EncodeToString(Sha1(key, msg)) }

func Sha224(key, msg []byte) []byte { return sum(sha256.New224, key, msg) }

func Sha224Hex(key, msg []byte) string { return hex.EncodeToString(Sha224(key, msg)) }

func Sha224Base64(key, msg []byte) string { return hex.EncodeToString(Sha224(key, msg)) }

func Sha256(key, msg []byte) []byte { return sum(sha256.New, key, msg) }

func Sha256Hex(key, msg []byte) string { return hex.EncodeToString(Sha256(key, msg)) }

func Sha256Base64(key, msg []byte) string { return hex.EncodeToString(Sha256(key, msg)) }

func Sha384(key, msg []byte) []byte { return sum(sha512.New384, key, msg) }

func Sha384Hex(key, msg []byte) string { return hex.EncodeToString(Sha384(key, msg)) }

func Sha384Base64(key, msg []byte) string { return hex.EncodeToString(Sha384(key, msg)) }

func Sha512(key, msg []byte) []byte { return sum(sha512.New, key, msg) }

func Sha512Hex(key, msg []byte) string { return hex.EncodeToString(Sha512(key, msg)) }

func Sha512Base64(key, msg []byte) string { return hex.EncodeToString(Sha512(key, msg)) }

func Sm3(key, msg []byte) []byte { return sum(sm3.New, key, msg) }

func Sm3Hex(key, msg []byte) string { return hex.EncodeToString(Sm3(key, msg)) }

func Sm3Base64(key, msg []byte) string { return hex.EncodeToString(Sm3(key, msg)) }

func sum(f func() hash.Hash, msg, key []byte) []byte {
	h := hmac.New(f, key)

	h.Write(msg)
	return h.Sum(nil)
}
