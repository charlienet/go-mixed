package bloom

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRedisStore(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.2.222:6379",
		Password: "123456",
	})

	store := newRedisStore(client, "abcdef", 10000)
	err := store.Set(1, 2, 3, 9, 1223)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(store.Test(1))
	t.Log(store.Test(1, 2, 3))
	t.Log(store.Test(4, 5, 8))
}
