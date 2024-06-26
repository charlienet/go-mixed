package bloom

import (
	"context"
	"math"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/hash"
	"github.com/charlienet/go-mixed/redis"
)

const DEFAULT_SIZE = 2 << 24

var seeds = []uint{7, 11, 13, 31, 37, 61, 79, 97}

type bitStore interface {
	Clear()
	Set(ctx context.Context, pos ...uint) error
	Test(pos ...uint) (bool, error)
}

type BloomFilter struct {
	bits  uint     // 布隆过滤器大小
	funcs uint     // 哈希函数数量
	store bitStore // 位图存储
}

type bloomOptions struct {
	redisClient redis.Client
	redisKey    string
}

type option func(*bloomOptions)

func WithRedis(redis redis.Client, key string) option {
	return func(bo *bloomOptions) {
		bo.redisClient = redis
		bo.redisKey = key
	}
}

// New 初始化布隆过滤器
// https://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html
func New(expectedInsertions uint, fpp float64, opts ...option) *BloomFilter {
	opt := &bloomOptions{}

	for _, f := range opts {
		f(opt)
	}

	bits := optimalNumOfBits(expectedInsertions, fpp)
	k := optimalNumOfHashFunctions(bits, expectedInsertions)
	if k > uint(len(seeds)) {
		k = uint(len(seeds))
	}

	bf := &BloomFilter{
		bits:  bits,
		funcs: k,
		store: createBitStore(opt, bits),
	}

	return bf
}

func (bf *BloomFilter) Add(ctx context.Context, data string) {
	offsets := bf.geOffsets([]byte(data))
	bf.store.Set(ctx, offsets...)
}

func (bf *BloomFilter) ExistString(data string) (bool, error) {
	return bf.Exists(bytesconv.StringToBytes(data))
}

func (bf *BloomFilter) Exists(data []byte) (bool, error) {
	if len(data) == 0 {
		return false, nil
	}

	offsets := bf.geOffsets(data)
	isSet, err := bf.store.Test(offsets...)
	if err != nil {
		return false, err
	}

	return isSet, nil
}

func (bf *BloomFilter) geOffsets(data []byte) []uint {
	offsets := make([]uint, bf.funcs)
	for i := uint(0); i < bf.funcs; i++ {
		offsets[i] = uint(hash.Murmur3(append(data, byte(seeds[i]))) % uint64(bf.bits))
	}

	return offsets
}

// 清空布隆过滤器
func (bf *BloomFilter) Clear() {
	bf.store.Clear()
}

func createBitStore(opt *bloomOptions, bits uint) bitStore {
	if opt.redisClient != nil {
		return newRedisStore(opt.redisClient, opt.redisKey, bits)
	}

	return newMemStore(bits)
}

// 计算优化的位图长度，
// n 期望放置元素数量，
// p 预期的误判概率
func optimalNumOfBits(n uint, p float64) uint {
	return (uint)(-float64(n) * math.Log(p) / (math.Log(2) * math.Log(2)))
}

// 计算哈希函数数量
func optimalNumOfHashFunctions(m, n uint) uint {
	return uint(math.Round(float64(m) / float64(n) * math.Log(2)))
}
