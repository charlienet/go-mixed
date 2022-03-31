package rand

import (
	"crypto/rand"
	"io"
	"sync"
	"time"

	"math/big"
	mrnd "math/rand"

	"github.com/charlienet/go-mixed/bytesconv"
)

const (
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	digit     = "0123456789"
	nomix     = "BCDFGHJKMPQRTVWXY2346789"
	letter    = uppercase + lowercase
	allChars  = uppercase + lowercase + digit
	hex       = digit + "ABCDEF"
	_         = allChars + "/+"
)

var (
	randSource mrnd.Source = mrnd.NewSource(time.Now().UnixNano())
	randLock   sync.Mutex
)

func init() {
	mrnd.Seed(time.Now().UnixNano())
}

type charScope struct {
	bytes   []byte
	length  int
	max     int
	bits    int
	mask    int
	lenFunc func(int) int
}

func StringScope(str string) *charScope {
	return strScope(str, nil)
}

func strScope(str string, f func(int) int) *charScope {
	len := len(str)

	scope := &charScope{
		bytes:   bytesconv.StringToBytes(str),
		length:  len,
		lenFunc: f,
		bits:    1,
	}

	for scope.mask < len {
		scope.bits++
		scope.mask = 1<<scope.bits - 1
	}

	scope.max = scope.mask / scope.bits

	return scope
}

var (
	Uppercase = StringScope(uppercase)                          // 大写字母
	Lowercase = StringScope(lowercase)                          // 小写字母
	Digit     = StringScope(digit)                              // 数字
	Nomix     = StringScope(nomix)                              // 不混淆字符
	Letter    = StringScope(letter)                             // 字母
	Hex       = strScope(hex, func(n int) int { return n * 2 }) // 十六进制字符
	AllChars  = StringScope(allChars)                           // 所有字符
)

// 生成指定长度的随机字符串
func (scope *charScope) Generate(length int) string {
	n := length
	if scope.lenFunc != nil {
		n = scope.lenFunc(n)
	}

	ret := make([]byte, n)
	for i, cache, remain := n-1, randInt63(), scope.max; i >= 0; {
		if remain == 0 {
			cache, remain = randInt63(), scope.max
		}

		if idx := int(cache & int64(scope.mask)); idx < scope.length {
			ret[i] = scope.bytes[idx]
			i--
		}

		cache >>= int64(scope.bits)
		remain--
	}

	return bytesconv.BytesToString(ret)
}

type scopeConstraint interface {
	~int | ~int32 | ~int64
}

// 生成区间 n >= 0, n < max
func Intn[T scopeConstraint](max T) T {
	n := mrnd.Int63n(int64(max))
	return T(n) % max
}

// 生成区间 n >= min, n < max
func IntRange[T scopeConstraint](min, max T) T {
	n := Intn(max - min)
	return T(n + min)
}

func CryptoRange[T scopeConstraint](min, max T) T {
	n := CryptoIntn(max - min)
	return min + n
}

func CryptoIntn[T ~int | ~int32 | ~int64](max T) T {
	b := big.NewInt(int64(max))
	n, _ := rand.Int(rand.Reader, b)

	return T(n.Int64())
}

func RandBytes(len int) ([]byte, error) {
	r := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, r)
	return r, err
}

func randInt63() int64 {
	var v int64

	randLock.Lock()
	v = randSource.Int63()
	randLock.Unlock()

	return v
}

func randNumber2(max int) int {
	return mrnd.Intn(max)
}
