package hash

import (
	"crypto/md5"
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

func Md5(msg []byte) hashResult { return sum(md5.New, msg) }

func Sha1(msg []byte) hashResult { return sum(sha1.New, msg) }

func Sha224(msg []byte) hashResult { return sum(sha256.New224, msg) }

func Sha256(msg []byte) hashResult { return sum(sha256.New, msg) }

func Sha384(msg []byte) hashResult { return sum(sha512.New384, msg) }

func Sha512(msg []byte) hashResult { return sum(sha512.New, msg) }

func Sm3(msg []byte) hashResult { return sum(sm3.New, msg) }

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

type hashResult []byte

func sum(f func() hash.Hash, msg []byte) hashResult {
	h := f()

	_, _ = h.Write(msg)
	return h.Sum(nil)
}

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
