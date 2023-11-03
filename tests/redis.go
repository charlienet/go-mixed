package tests

import (
	"context"
	"log"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/charlienet/go-mixed/redis"
	"github.com/stretchr/testify/assert"
)

var DefaultRedis = redis.RedisOption{
	Addr:     "redis:6379",
	Password: "123456",
}

func RunOnDefaultRedis(t assert.TestingT, fn func(rdb redis.Client)) {
	RunOnRedis(t, fn, DefaultRedis)
}

func RunOnRedis(t assert.TestingT, fn func(rdb redis.Client), opt ...redis.RedisOption) {
	var redis redis.Client
	var clean func()
	var err error

	redis, clean, err = CreateRedis(opt...)
	assert.Nil(t, err, err)

	defer clean()
	fn(redis)
}

func CreateRedis(opt ...redis.RedisOption) (r redis.Client, clean func(), err error) {
	if len(opt) > 0 {
		return createRedisClient(opt[0])
	} else {
		return createMiniRedis()
	}
}

func createRedisClient(opt redis.RedisOption) (r redis.Client, clean func(), err error) {
	rdb := redis.New(&opt)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, nil, err
	}

	return rdb, func() { rdb.Close() }, nil
}

func createMiniRedis() (r redis.Client, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	addr := mr.Addr()
	log.Println("mini redis run at:", addr)

	rdb := redis.New(&redis.RedisOption{
		Addrs: []string{addr},
	})

	return rdb, func() {
		ch := make(chan struct{})

		go func() {
			rdb.Close()
			mr.Close()
			close(ch)
		}()

		select {
		case <-ch:
		case <-time.After(time.Second * 5):
		}
	}, nil
}
