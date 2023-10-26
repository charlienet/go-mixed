package idgenerator

import (
	"testing"

	"github.com/charlienet/go-mixed/idGenerator/store"
	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
)

func TestBufferAlloc(t *testing.T) {

	tests.RunOnRedis(t, func(rdb redis.Client) {
		f := func() (*store.Segment, error) {
			return store.NewRedisStore("sss", rdb).Assign(3, 99, 10)
		}

		b := newDoubleBuffer(f)

		for i := 0; i < 80; i++ {
			t.Log(b.allot())
		}

	}, redis.ReidsOption{Addr: "192.168.123.50:6379", Password: "123456"})
}
