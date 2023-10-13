package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/charlienet/go-mixed/redis"
	"github.com/stretchr/testify/assert"
)

func RunOnSpecifiedRedis(t *testing.T, fn func(client redis.Client), addr ...string) {
	rdb := redis.New(&redis.ReidsOption{
		Addrs: addr,
	})
	defer rdb.Close()

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		t.Fatal(err)
	}

	fn(rdb)
}

func RunOnRedis(t *testing.T, fn func(client redis.Client)) {
	redis, clean, err := createMiniRedis()
	assert.Nil(t, err)

	defer clean()

	fn(redis)
}

func createMiniRedis() (r redis.Client, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	addr := mr.Addr()
	log.Println("mini redis run at:", addr)

	return redis.New(&redis.ReidsOption{
			Addrs: []string{addr},
		}), func() {
			ch := make(chan struct{})

			go func() {
				mr.Close()
				close(ch)
			}()

			select {
			case <-ch:
			case <-time.After(time.Second):
			}
		}, nil
}
