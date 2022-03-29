package hash

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"hash/fnv"

	"github.com/cespare/xxhash/v2"
	"github.com/spaolacci/murmur3"
	"github.com/tjfoc/gmsm/sm3"
)

func Sha1(msg []byte) []byte { return sum(sha1.New, msg) }

func Sha1Hex(msg []byte) string { return hex.EncodeToString(Sha1(msg)) }

func Sha1Base64(msg []byte) string { return base64.StdEncoding.EncodeToString(Sha1(msg)) }

func Sha224(msg []byte) []byte { return sum(sha256.New224, msg) }

func Sha224Hex(msg []byte) string { return hex.EncodeToString(Sha224(msg)) }

func Sha224Base64(msg []byte) string { return base64.StdEncoding.EncodeToString(Sha224(msg)) }

func Sha256(msg []byte) []byte { return sum(sha256.New, msg) }

func Sha256Hex(msg []byte) string { return hex.EncodeToString(Sha256(msg)) }

func Sha256Base64(msg []byte) string { return base64.StdEncoding.EncodeToString(Sha256(msg)) }

func Sha384(msg []byte) []byte { return sum(sha512.New384, msg) }

func Sha384Hex(msg []byte) string { return hex.EncodeToString(Sha384(msg)) }

func Sha384Base64(msg []byte) string { return base64.StdEncoding.EncodeToString(Sha384(msg)) }

func Sha512(msg []byte) []byte { return sum(sha512.New, msg) }

func Sha512Hex(msg []byte) string { return hex.EncodeToString(Sha512(msg)) }

func Sha512Base64(msg []byte) string { return base64.StdEncoding.EncodeToString(Sha512(msg)) }

func Sm3(msg []byte) []byte { return sum(sm3.New, msg) }

func Sm3Hex(msg []byte) string { return hex.EncodeToString(Sm3(msg)) }

func Sm3Base64(msg []byte) string { return base64.StdEncoding.EncodeToString(Sm3(msg)) }

func Murmur3(msg []byte) uint64 {
	return murmur3.Sum64(msg)
}

func XXhash(msg []byte) []byte {
	d := xxhash.New()
	_, _ = d.Write(msg)
	return d.Sum(nil)
}

func XXHashUint64(msg []byte) uint64 {
	h := xxhash.New()
	_, _ = h.Write(msg)
	return h.Sum64()
}

func Funv32(msg []byte) uint32 {
	h := fnv.New32()
	_, _ = h.Write(msg)
	return h.Sum32()
}

func Funv64(msg []byte) uint64 {
	h := fnv.New64()
	_, _ = h.Write(msg)
	return h.Sum64()
}

func sum(f func() hash.Hash, msg []byte) []byte {
	h := f()

	_, _ = h.Write(msg)
	return h.Sum(nil)
}
