package bloom

import "github.com/bits-and-blooms/bitset"

const DEFAULT_SIZE = 2 << 24

var seeds = []uint{7, 11, 13, 31, 37, 61}

type simplehash struct {
	cap  uint
	seed uint
}

type BloomFilter struct {
	set   *bitset.BitSet
	funcs [6]simplehash
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(bf.funcs); i++ {
		bf.funcs[i] = simplehash{DEFAULT_SIZE, seeds[i]}
	}
	bf.set = bitset.New(DEFAULT_SIZE)
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

func (s simplehash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*s.seed + uint(value[i])
	}
	return (s.cap - 1) & result
}
