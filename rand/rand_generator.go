package rand

import (
	mrnd "math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	seed  = time.Now().UnixNano()                 // 随机数种子
	souce = mrnd.NewSource(time.Now().UnixNano()) // 用于初始化的随机数生成器
)

type randGenerator struct {
	source mrnd.Source
	r      uint32
}

func NewRandGenerator() *randGenerator {
	return &randGenerator{
		source: mrnd.NewSource(getSeed()),
	}
}

func getSeed() int64 {
	seed ^= souce.Int63()
	return seed
}

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r *randGenerator) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int63() & (n - 1)
	}
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > max {
		v = r.Int63()
	}

	return v % n
}

func (r *randGenerator) Int31() int32 {
	return int32(r.Int63() >> 32)
}

func (r *randGenerator) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}

	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > max {
		v = r.Int31()
	}
	return v % n
}

func (r *randGenerator) Int() int {
	u := uint(r.Int63())
	return int(u << 1 >> 1) // clear sign bit if int == int32
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (r *randGenerator) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

func (r *randGenerator) Int63() int64 {
	r.lock()
	i := r.source.Int63()
	r.unlock()

	return i
}

func (g *randGenerator) lock() {
	for !atomic.CompareAndSwapUint32(&g.r, 0, 1) {
		runtime.Gosched()
	}
}

func (g *randGenerator) unlock() {
	atomic.StoreUint32(&g.r, 0)
}
