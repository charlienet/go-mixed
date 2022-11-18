package redis

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetSet(t *testing.T) {
	runOnRedis(t, func(client *Redis) {
		ctx := context.Background()

		val, err := client.GetSet(ctx, "hello", "world")
		assert.NotNil(t, err)
		assert.Equal(t, "", val)

		val, err = client.Get(ctx, "hello")
		assert.Nil(t, err)
		assert.Equal(t, "world", val)

		val, err = client.GetSet(ctx, "hello", "newworld")
		assert.Nil(t, err)
		assert.Equal(t, "world", val)

		val, err = client.Get(ctx, "hello")
		assert.Nil(t, err)
		assert.Equal(t, "newworld", val)

		ret, err := client.Del(ctx, "hello")
		assert.Nil(t, err)
		assert.Equal(t, 1, ret)
	})
}

func TestRedis_SetGetDel(t *testing.T) {
	runOnRedis(t, func(client *Redis) {
		ctx := context.Background()

		err := client.Set(ctx, "hello", "world")
		assert.Nil(t, err)

		val, err := client.Get(ctx, "hello")
		assert.Nil(t, err)
		assert.Equal(t, "world", val)
		ret, err := client.Del(ctx, "hello")
		assert.Nil(t, err)
		assert.Equal(t, 1, ret)
	})
}

func runOnRedis(t *testing.T, fn func(client *Redis)) {
	redis, clean, err := CreateMiniRedis()
	assert.Nil(t, err)

	defer clean()

	fn(redis)
}

func CreateMiniRedis() (r *Redis, clean func(), err error) {
	mr, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}

	addr := mr.Addr()
	log.Println("mini redis run at:", addr)

	return New(addr), func() {
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
