package redis_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetSet(t *testing.T) {
	tests.RunOnRedis(t, func(client redis.Client) {
		ctx := context.Background()

		val, err := client.GetSet(ctx, "hello", "world").Result()
		assert.NotNil(t, err)
		assert.Equal(t, "", val)

		val, err = client.Get(ctx, "hello").Result()
		assert.Nil(t, err)
		assert.Equal(t, "world", val)

		val, err = client.GetSet(ctx, "hello", "newworld").Result()
		assert.Nil(t, err)
		assert.Equal(t, "world", val)

		val, err = client.Get(ctx, "hello").Result()
		assert.Nil(t, err)
		assert.Equal(t, "newworld", val)

		ret, err := client.Del(ctx, "hello").Result()
		assert.Nil(t, err)
		assert.Equal(t, 1, ret)
	})
}

func TestRedis_SetGetDel(t *testing.T) {
	tests.RunOnRedis(t, func(client redis.Client) {
		ctx := context.Background()

		_, err := client.Set(ctx, "hello", "world", 0).Result()
		assert.Nil(t, err)

		val, err := client.Get(ctx, "hello").Result()
		assert.Nil(t, err)
		assert.Equal(t, "world", val)
		ret, err := client.Del(ctx, "hello").Result()
		assert.Nil(t, err)
		assert.Equal(t, int64(1), ret)
	})
}

func TestPubSub(t *testing.T) {
	tests.RunOnRedis(t, func(client redis.Client) {
		ctx := context.Background()

		c := "chat"
		quit := false

		total := 0
		mu := &sync.Mutex{}
		f := func(wg *sync.WaitGroup) {
			wg.Add(1)
			var receivedCount int = 0

			sub := client.Subscribe(ctx, c)
			defer sub.Close()

			for {
				select {
				case <-sub.Channel():
					receivedCount++
				// case <-quit:

				default:
					if quit {
						mu.Lock()
						total += receivedCount
						mu.Unlock()

						t.Logf("Subscriber received %d message  %d", receivedCount, total)
						wg.Done()

						return
					}
				}
			}

			// for msg := range sub.Channel() {
			// 	if strings.EqualFold(msg.Payload, "quit") {
			// 		break
			// 	}

			// 	receivedCount++
			// }

		}

		var wg = &sync.WaitGroup{}
		go f(wg)
		go f(wg)
		go f(wg)

		for i := 0; i < 20000; i++ {

			n, err := client.Publish(ctx, c, fmt.Sprintf("hello %d", i)).Result()
			if err != nil {
				t.Log(err)
			}

			_ = n

			// t.Logf("%d clients received the message\n", n)
		}

		// for i := 0; i < 20; i++ {
		// 	client.Publish(ctx, c, "quit")
		// }

		t.Log("finished send message")

		time.Sleep(time.Second * 5)
		quit = true

		wg.Wait()

		time.Sleep(time.Second * 2)
		t.Logf("total received %d message", total)
	})
}
