package locker_test

import (
	"testing"

	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
)

func TestRedisDistrbutedLocker(t *testing.T) {
	tests.RunOnDefaultRedis(t, func(rdb redis.Client) {
	})
}
