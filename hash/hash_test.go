package hash_test

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"

	"github.com/charlienet/go-mixed/hash"
	"github.com/charlienet/go-mixed/rand"
	"github.com/stretchr/testify/assert"
)

func TestHashComplie(t *testing.T) {
	abc, err := hash.New("MD5")
	if err != nil {

	}
	b, _ := hex.DecodeString(rand.Hex.Generate(16))
	assert.False(t, abc.Verify([]byte("source"), b))
}

func TestEncode(t *testing.T) {
	t.Log(hash.Sha1([]byte{0x31}).Base64())
	t.Log(hash.Sha1([]byte{0x31}).Hex())
}

func TestXXHash(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(hex.EncodeToString(hash.XXhash([]byte(strconv.Itoa(i)))), "  ", hash.XXHashUint64([]byte(strconv.Itoa(i))))
	}
}

func TestMurmur3(t *testing.T) {
	t.Log(hash.Murmur3([]byte("123")))
	t.Log(hash.XXHashUint64([]byte("123")))
}

func TestFnv(t *testing.T) {
	for i := 0; i < 100; i++ {
		bytes := []byte(fmt.Sprintf("%d", i))
		t.Log(hash.Funv32(bytes))
	}
}

func BenchmarkHash(b *testing.B) {
	bytes := []byte("abcdefdg")
	b.Run("xxhash", func(b *testing.B) {
		doBenchmark(func() {
			hash.XXHashUint64(bytes)
		}, b)
	})

	b.Run("murmur3", func(b *testing.B) {
		doBenchmark(func() {
			hash.Murmur3(bytes)
		}, b)
	})

	b.Run("fnv", func(b *testing.B) {
		doBenchmark(func() {
			hash.Funv64(bytes)
		}, b)
	})

	b.Run("sm3", func(b *testing.B) {
		doBenchmark(func() {
			hash.Sm3(bytes)
		}, b)
	})

	b.Run("md5", func(b *testing.B) {
		doBenchmark(func() {
			hash.Md5(bytes)
		}, b)
	})

	b.Run("sha256", func(b *testing.B) {
		doBenchmark(func() {
			hash.Sha256(bytes)
		}, b)
	})

}

func doBenchmark(f func(), b *testing.B) {
	for i := 0; i < b.N; i++ {
		f()
	}
}
