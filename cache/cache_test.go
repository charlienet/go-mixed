package cache

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	defaultKey = "u-000"
)

func TestNewCache(t *testing.T) {
	// c, err := NewCacheBuilder().
	// 	UseRedis(RedisConfig{
	// 		Addrs:    []string{"192.168.2.222:6379"},
	// 		Password: "123456",
	// 	}).
	// 	WithPrefix("cache_test").
	// 	WithLogger(logx.NewLogrus()).
	// 	Build()

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// ctx := context.Background()
	// c.Set(ctx, "abc", "value", time.Minute*10)

	// var s string
	// c.Get(ctx, "abc", &s)

	// t.Log(s)
}

type SimpleUser struct {
	FirstName string
	LastName  string
}

func TestMemCache(t *testing.T) {
	// b, _ := bigcache.NewBigCache(bigcache.BigCacheConfig{})
	// var mems = []MemCache{
	// 	NewFreeCache(10 * 1024 * 1024),
	// 	b,
	// }

	// u := SimpleUser{FirstName: "Radomir", LastName: "Sohlich"}
	// encoded, _ := bytesconv.Encode(u)
	// for _, m := range mems {
	// 	m.Set(defaultKey, encoded, time.Second)
	// 	ret, _ := m.Get(defaultKey)

	// 	var u2 SimpleUser
	// 	bytesconv.Decode(ret, &u2)
	// 	t.Log(u2)
	// }
}

func TestDistributedCache(t *testing.T) {
	c := NewRedis(RedisConfig{Addrs: []string{"192.168.2.222:6379"}, DB: 6, Password: "123456", Prefix: "abcdef"})

	ctx := context.Background()
	if err := c.Ping(ctx); err != nil {
		t.Fatal(err)
	}

	t.Log(c.Exist(ctx, defaultKey))

	u := SimpleUser{FirstName: "redis client"}

	var u2 SimpleUser
	c.Get(ctx, defaultKey, &u2)

	c.Set(ctx, defaultKey, u, time.Minute*10)
	t.Log(c.Exist(ctx, defaultKey))

	if err := c.Get(ctx, defaultKey, &u2); err != nil {
		t.Fatal("err:", err)
	}
	t.Logf("%+v", u2)

	// c.Delete(defaultKey)
}

func TestGetFn(t *testing.T) {
	c := buildCache()
	var u2 SimpleUser

	c.GetFn(context.Background(), defaultKey, &u2, func(ctx context.Context) (out any, err error) {
		v := &u2
		v.FirstName = "abc"
		v.LastName = "aaaa"

		return nil, nil
	}, time.Minute*1)

	t.Logf("%+v", u2)
}

func TestGetFromSource(t *testing.T) {
	var count int32

	n := 10
	c := &Cache{}
	wg := &sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			c.getFromSource(context.Background(), defaultKey, func(ctx context.Context) (any, error) {
				atomic.AddInt32(&count, 1)
				time.Sleep(time.Second)

				return "abc", nil
			})

			wg.Done()
		}()
	}

	wg.Wait()
	t.Log("count:", count)
}

func BenchmarkMemCache(b *testing.B) {
}

func load() (any, error) {
	return nil, nil
}

func buildCache() *Cache {
	c, err := New(
		WithFreeCache(10*1024*1024),
		WithRedis(RedisConfig{Addrs: []string{"192.168.2.222:6379"}, DB: 6, Password: "123456"}))

	if err != nil {
		panic(err)
	}

	return c
}
