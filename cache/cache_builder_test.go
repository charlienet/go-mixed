package cache

import (
	"testing"
	"time"

	"github.com/charlienet/go-mixed/logx"
)

func TestBuilder(t *testing.T) {
	cache, err := NewCacheBuilder().
		WithLogger(logx.NewLogrus(logx.WithFormatter(logx.NewNestedFormatter(logx.NestedFormatterOption{
			Color: true,
		})))).
		WithRedis(RedisConfig{
			Addrs:    []string{"192.168.2.222:6379"},
			Password: "123456",
		}).
		WithBigCache(BigCacheConfig{}).
		// WithFreeCache(10 * 1024 * 1024).
		Build()

	if err != nil {
		t.Fatal(err)
	}

	u := SimpleUser{FirstName: "Radomir", LastName: "Sohlich"}
	t.Log(cache.Set(defaultKey, u, time.Minute*10))
}
