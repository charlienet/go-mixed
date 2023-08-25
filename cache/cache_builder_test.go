package cache

import (
	"testing"
	"time"
)

func TestBuilder(t *testing.T) {
	now := time.Now()
	t.Log(now)
	t1, _ := time.ParseDuration("9h27m")
	t1 += time.Hour * 24
	t2, _ := time.ParseDuration("16h28m")
	t.Log(t1)
	t.Log(t2)

	f := time.Date(2022, time.December, 12, 8, 0, 0, 0, time.Local)
	t.Log(f.Sub(time.Now()))

	// cache, err := New(
	// 	WithLogger(logx.NewLogrus(logx.WithNestedFormatter(logx.NestedFormatterOption{
	// 		Color: true,
	// 	}))).
	// 	UseRedis(RedisConfig{
	// 		Addrs:    []string{"192.168.2.222:6379"},
	// 		Password: "123456",
	// 	}).
	// 	UseBigCache(bigcache.BigCacheConfig{}).
	// 	Build()

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()

	// u := SimpleUser{FirstName: "Radomir", LastName: "Sohlich"}
	// t.Log(cache.Set(ctx, defaultKey, u, time.Minute*10))
}
