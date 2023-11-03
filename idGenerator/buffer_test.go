package idgenerator

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/charlienet/go-mixed/idGenerator/store"
	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
)

func TestBufferAlloc(t *testing.T) {

	tests.RunOnDefaultRedis(t, func(rdb redis.Client) {
		f := func() (*store.Segment, error) {
			return store.NewRedisStore("sss", rdb).Assign(3, 99, 10)
		}

		b := newDoubleBuffer(f)

		for i := 0; i < 80; i++ {
			t.Log(b.allot())
		}

	})
}

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			println("协程退出", ctx.Err().Error())
		case <-time.After(time.Second * 100):
			println("协程超时")
		}
	}()

	wg.Wait()

	println("应用退出")
}
