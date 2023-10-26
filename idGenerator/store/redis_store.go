package store

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/charlienet/go-mixed/rand"
	"github.com/charlienet/go-mixed/redis"
)

const (
	allocatedKey = "allocated"
	sequenceKey  = "sequence"
	delaySecond  = 10
)

const (
	// 机器码分配及保活，检查机器码键值是否匹配，如匹配成功对键进行延时。如不匹配需要重新申请机器码
	// 请求参数  机器码，机器识别码，机器码最大值
	machineLua = `
	local newKey = KEYS[1]..":"..ARGV[1]
	if redis.call("GET", newKey) == ARGV[2] then
		redis.call("EXPIRE", newKey, 10)
		return tonumber(ARGV[1])
	else
		for i = 0, ARGV[3], 1 do
			newKey = KEYS[1]..":"..tostring(i)

			if redis.call("EXISTS", newKey) == 0 then
				redis.call("SET", newKey, ARGV[2], "EX", 10)
				return i
			end
		end
		return -1
	end`

	// 序列分配min, max, step
	segmentLua = `
local key = KEYS[1]
local min = tonumber(ARGV[1])
local max = tonumber(ARGV[2])
local step = tonumber(ARGV[3])

if step > max then
	step = max
end

if redis.call("EXISTS", key) == 0 then
	redis.call("SET", key, step)
	return {min, step, 0}
end

local begin = tonumber(redis.call("GET", key))
local increase = redis.call("INCRBY", key, step)
local reback = 0

if begin >= max then
	begin = min
	increase = step

	redis.call("SET", key, step)
	reback = 1
end

if increase > max then
	increase = max
	redis.call("SET", key, increase)
end

return {begin, increase, reback}
`
)

type redisStore struct {
	rdb         redis.Client
	machinekey  string        // 机器码键
	sequenceKey string        // 序列键
	value       string        // 随机键值
	machineCode int64         // 机器码
	close       chan struct{} // 关闭保活协程
	isRunning   bool          // 是否已经关闭
	mu          sync.Mutex
}

func NewRedisStore(key string, rdb redis.Client) *redisStore {
	return &redisStore{
		rdb:         rdb,
		machinekey:  rdb.JoinKeys(key, allocatedKey),
		sequenceKey: rdb.JoinKeys(key, sequenceKey),
		value:       rand.Hex.Generate(24),
		close:       make(chan struct{}),
	}
}

// 分配机器标识，分配值为-1时表示分配失败
func (s *redisStore) UpdateMachineCode(max int64) (int64, error) {
	err := s.updateMachine(max)
	if err != nil {
		return -1, err
	}

	// 关闭原协程，开启新的保活协程

	// if s.isRunning {
	// 	s.close <- struct{}{}
	// }

	// if !s.isRunning {
	// s.close <- struct{}{}
	go s.keepAlive(max)
	// }

	return s.machineCode, nil
}

func (s *redisStore) MachineCode() int64 {
	return s.machineCode
}

func (s *redisStore) Assign(min, max, step int64) (*Segment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// 序列段分配
	key := s.rdb.JoinKeys(s.sequenceKey, fmt.Sprintf("%v", s.machineCode))
	r, err := s.rdb.Eval(ctx, segmentLua, []string{key}, min, max, step).Result()
	if err != nil {
		return &Segment{}, err
	}

	start, end, reback := split(r)
	return &Segment{start: start, end: end, current: start, reback: reback}, err
}

func split(r any) (start, end int64, reback bool) {
	if result, ok := r.([]any); ok {
		start = result[0].(int64)
		end = result[1].(int64)
		reback = result[2].(int64) == 1
	}

	return
}

func (s *redisStore) updateMachine(max int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := s.rdb.Eval(ctx, machineLua, []string{s.machinekey}, s.machineCode, s.value, max).Result()
	if err != nil {
		return err
	}

	if r == nil {
		return errors.New("failed to obtain machine code")
	}

	s.machineCode = r.(int64)
	if s.machineCode == -1 {
		return errors.New("machine code allocation failed")
	}

	return nil
}

func (s *redisStore) Close() {
	s.close <- struct{}{}
	s.isRunning = false
}

func (s *redisStore) keepAlive(max int64) {
	t := time.NewTicker(time.Second * (delaySecond / 3))
	defer t.Stop()

	for {
		select {
		case <-t.C:
			// println("当前机器码:", s.machineCode)
			err := s.updateMachine(max)
			if err != nil {
				fmt.Println("err:", err.Error())
			}
		case <-s.close:
			println("保活停止")
			return
		}
	}
}
