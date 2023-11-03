package store

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "embed"

	"github.com/charlienet/go-mixed/rand"
	"github.com/charlienet/go-mixed/redis"
)

//go:embed redis_id_store.lua
var redis_id_function string

var once sync.Once

type redisStore struct {
	rdb         redis.Client
	key         string        // 缓存键
	machine     string        // 随机键值(机器标识)
	machineCode int64         // 机器码
	max         int64         // 机器码的最大值
	close       chan struct{} // 关闭保活协程
	isRunning   bool          // 是否已经关闭
	mu          sync.Mutex
}

func NewRedisStore(key string, rdb redis.Client) *redisStore {
	once.Do(func() { rdb.LoadFunction(redis_id_function) })

	return &redisStore{
		rdb:         rdb,
		key:         key,
		machineCode: -1,
		machine:     rand.Hex.Generate(24),
		close:       make(chan struct{}),
	}
}

// 分配机器标识，分配值为-1时表示分配失败
func (s *redisStore) UpdateMachineCode(max int64) (int64, error) {
	s.max = max

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

	// 序列段分配 机器码，机器标识，步长，序列最小值，序列最大值，机器码最大值
	r, err := s.rdb.FCall(ctx, "allocateSerial", []string{s.key}, s.machineCode, s.machine, step, min, max, s.max).Result()
	if err != nil {
		return &Segment{}, err
	}

	machineCode, start, end, reback := split(r)
	s.machineCode = machineCode

	return &Segment{start: start, end: end, current: start, reback: reback}, err
}

func split(r any) (machineCode, start, end int64, reback bool) {
	if result, ok := r.([]any); ok {
		machineCode = result[0].(int64)
		start = result[1].(int64)
		end = result[2].(int64)
		reback = result[3].(int64) == 1
	}

	return
}

func (s *redisStore) updateMachine(max int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//机器码当前值，机器标识，机器码最大值
	r, err := s.rdb.FCall(ctx, "updateMachineCode", []string{s.key}, s.machineCode, s.machine, s.max).Result()
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
	t := time.NewTicker(time.Second * 2)
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
