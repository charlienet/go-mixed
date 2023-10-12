package bloom

import (
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/charlienet/go-mixed/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedisStore(t *testing.T) {
	runOnRedis(t, func(client redis.Client) {
		store := newRedisStore(client, "abcdef", 10000)
		err := store.Set(1, 2, 3, 9, 1223)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(store.Test(1))
		t.Log(store.Test(1, 2, 3))
		t.Log(store.Test(4, 5, 8))
	})
}

func runOnRedis(t *testing.T, fn func(client redis.Client)) {
	redis, clean, err := CreateMiniRedis()
	assert.Nil(t, err)

	defer clean()

	fn(redis)
}

func CreateMiniRedis() (r redis.Client, clean func(), err error) {
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
