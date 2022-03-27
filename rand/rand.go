package rand

import (
	"crypto/rand"
	"io"
	"sync"
	"time"

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

var randSource mrnd.Source = mrnd.NewSource(time.Now().UnixNano())
var randLock sync.Mutex

// 生成指定长度的随机字符串
func (scope *charScope) RandString(length int) string {
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

// 获取指定范围内的随机数
func RandInt(min, max int) int {
	n := randNumber2(max - min)
	return n + min
}

// 生成指定范围的随机数
func RandInt32(min, max int32) int32 {
	return int32(RandInt(int(min), int(max)))
}

func RandBytes(len int) ([]byte, error) {
	r := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, r)
	return r, err
}

// func randNumber(max int64) (int64, error) {
// 	r, err := rand.Int(rand.Reader, big.NewInt(max+1))
// 	return r.Int64(), err
// }

func randInt63() int64 {
	var v int64

	randLock.Lock()
	v = randSource.Int63()
	randLock.Unlock()

	return v
}

func randNumber2(max int) int {
	rnd := mrnd.New(randSource)
	return rnd.Intn(max)
}
