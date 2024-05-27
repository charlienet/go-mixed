package redis

import (
	"context"
	"testing"
	"time"

	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
)

func TestCreateRedisStore(t *testing.T) {
	tests.RunOnDefaultRedis(t, func(rdb redis.Client) {
		keyName := "source"

		l := NewRedisStore("locker_key", rdb)
		ret := l.TryLock(context.Background(), keyName)
		if !ret {
			t.Log("加锁失败")
		}

		l.Lock(context.Background(), keyName)
		t.Log("锁重入完成")

		l.Unlock(context.Background(), keyName)

		time.Sleep(time.Second * 15)

		// l.Unlock(context.Background())
	})
}
