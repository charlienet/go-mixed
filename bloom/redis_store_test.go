package bloom

import (
	"testing"

	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
)

func TestRedisStore(t *testing.T) {
	tests.RunOnRedis(t, func(client redis.Client) {
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
