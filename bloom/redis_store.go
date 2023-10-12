package bloom

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/charlienet/go-mixed/redis"
)

const (
	// ARGV:偏移量offset数组
	// KYES[1]: setbit操作的key
	// 全部设置为1
	setScript = `
		for _, offset in ipairs(ARGV) do
			redis.call("setbit", KEYS[1], offset, 1)
		end
    `

	//ARGV:偏移量offset数组
	//KYES[1]: setbit操作的key
	//检查是否全部为1
	testScript = `
	for _, offset in ipairs(ARGV) do
		if tonumber(redis.call("getbit", KEYS[1], offset)) == 0 then
			return false
		end
	end
	return true
    `
)

var ErrTooLargeOffset = errors.New("超出最大偏移量")

var _ bitStore = &redisBitSet{}

// 使用Redis存储位图
type redisBitSet struct {
	store redis.Client
	key   string
	bits  uint
}

func newRedisStore(store redis.Client, key string, bits uint) *redisBitSet {
	return &redisBitSet{
		store: store,
		key:   key,
		bits:  bits,
	}
}

func (s *redisBitSet) Set(offsets ...uint) error {
	args, err := s.buildOffsetArgs(offsets)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	_, err = s.store.Eval(ctx, setScript, []string{s.key}, args).Result()

	//底层使用的是go-redis,redis.Nil表示操作的key不存在
	//需要针对key不存在的情况特殊判断
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func (s *redisBitSet) Test(offsets ...uint) (bool, error) {
	args, err := s.buildOffsetArgs(offsets)
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	resp, err := s.store.Eval(ctx, testScript, []string{s.key}, args).Result()

	// key 不存在，表示还未存放任何数据
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	exists, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return exists == 1, nil
}

func (s *redisBitSet) Clear() {

}

func (r *redisBitSet) buildOffsetArgs(offsets []uint) ([]string, error) {
	args := make([]string, 0, len(offsets))
	for _, offset := range offsets {
		if offset >= r.bits {
			return nil, ErrTooLargeOffset
		}

		args = append(args, strconv.FormatUint(uint64(offset), 10))
	}
	return args, nil
}
