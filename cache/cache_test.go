package cache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	r := NewRedis(&RedisConfig{})
	if err := r.Ping(); err != nil {
		t.Fatal(err)
	}

	c, err := NewCache(
		WithDistributdCache(r),
		WithPrefix("cache_test"))
	if err != nil {
		t.Fatal(err)
	}

	c.Set("abc", "value", time.Minute*10)
}

type SimpleUser struct {
	FirstName string
	LastName  string
}

func TestMem(t *testing.T) {
	c, err := NewCache(WithFreeCache(10 * 1024 * 1024))
	if err != nil {
		t.Fatal(err)
	}
	key := "u-000"
	u := SimpleUser{FirstName: "Radomir", LastName: "Sohlich"}

	c.Set(key, u, time.Second)

	var u2 SimpleUser
	c.Get(key, &u2)

	t.Logf("%+v", u2)
}

func TestDistributedCache(t *testing.T) {
	key := "key-001"
	c := NewRedis(&RedisConfig{Addrs: []string{"192.168.2.222:6379"}, DB: 6, Password: "123456"})

	if err := c.Ping(); err != nil {
		t.Fatal(err)
	}
	u := SimpleUser{FirstName: "redis client"}
	c.Set(key, u, time.Second)

	var u2 SimpleUser
	if err := c.Get(key, &u2); err != nil {
		t.Fatal("err:", err)
	}
	t.Logf("%+v", u2)
}

func TestGetFn(t *testing.T) {
	c, err := NewCache(WithBigCache(&BigCacheConfig{}))
	if err != nil {
		t.Fatal(err)
	}
	key := "u-000"

	var u2 SimpleUser
	c.GetFn(key, &u2, func() (out any, err error) {
		v := &u2
		v.FirstName = "abc"

		return nil, nil
	}, time.Second)

	t.Logf("%+v", u2)
}

func BenchmarkMemCache(b *testing.B) {
}

func load() (any, error) {
	return nil, nil
}
