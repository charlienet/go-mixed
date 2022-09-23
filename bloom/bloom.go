package bloom

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/charlienet/go-mixed/locker"
)

const DEFAULT_SIZE = 2 << 24

var seeds = []uint{7, 11, 13, 31, 37, 61}

type simplehash struct {
	cap  uint
	seed uint
}

type BloomFilter struct {
	size  int            // 布隆过滤器大小
	set   *bitset.BitSet // 位图
	funcs [6]simplehash  // 哈希函数
	lock  locker.RWLocker
}

type bloomOptions struct {
	Size int
}

type option func(*bloomOptions)

// 布隆过滤器中所有位长度，请根据存储数量进行评估
func WithSize(size int) option {
	return func(bo *bloomOptions) {
		bo.Size = size
	}
}

func NewBloomFilter(opts ...option) *BloomFilter {
	opt := &bloomOptions{
		Size: DEFAULT_SIZE,
	}

	for _, f := range opts {
		f(opt)
	}

	bf := &BloomFilter{
		size: opt.Size,
		lock: locker.NewRWLocker(),
	}

	for i := 0; i < len(bf.funcs); i++ {
		bf.funcs[i] = simplehash{uint(opt.Size), seeds[i]}
	}
	bf.set = bitset.New(uint(opt.Size))
	return bf
}

func (bf *BloomFilter) Add(value string) {
	funcs := bf.funcs[:]

	for _, f := range funcs {
		bf.set.Set(f.hash(value))
	}

}

func (bf *BloomFilter) Contains(value string) bool {
	if value == "" {
		return false
	}
	ret := true

	funcs := bf.funcs[:]
	for _, f := range funcs {
		ret = ret && bf.set.Test(f.hash(value))
	}

	return ret
}

// 清空布隆过滤器
func (bf *BloomFilter) Clear() {
	bf.set.ClearAll()
}

func (s simplehash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*s.seed + uint(value[i])
	}
	return (s.cap - 1) & result
}
