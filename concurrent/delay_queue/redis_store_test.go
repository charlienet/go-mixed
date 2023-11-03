package delayqueue

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
	"github.com/stretchr/testify/assert"
)

const (
	redisAddr      = "192.168.123.100:6379"
	delay_queue    = "delay_queue"
	execute_queue  = "execute_queue"
	delay_task_set = "task_set"
)

func TestRedis(t *testing.T) {
	tests.RunOnRedis(t, func(client redis.Client) {
		defer client.Close()

		q := New[delayTask]().UseRedis(delay_queue, execute_queue, delay_task_set, client)

		err := q.Push(delayTask{
			Message: "abc1111111111111",
			At:      time.Now().Add(time.Second * 2)})

		if err != nil {
			t.Fatal(err)
		}

		t.Log(time.Now())

		task, _ := q.Pop()
		t.Logf("%+v", task)

		t.Log(time.Now())

		task.execute()

	}, redis.RedisOption{Addr: redisAddr, Prefix: "redis_test"})
}

func TestMutiTask(t *testing.T) {
	tests.RunOnRedis(t, func(client redis.Client) {
		defer client.Close()

		timer := time.NewTimer(time.Second)
		ticker := time.NewTicker(time.Second)

		timer.Reset(time.Microsecond)
		ticker.Reset(time.Millisecond)

		store := newRedisStroe[delayTask](delay_queue, execute_queue, delay_task_set, client)

		for i := 1; i <= 5; i++ {
			store.Push(context.Background(), delayTask{
				Message: fmt.Sprintf("abc:%d", i),
				At:      time.Now().Add(time.Second * time.Duration(i)),
			})
		}

		for !store.IsEmpty() {
			v, err := store.Pop()
			assert.Nil(t, err)
			t.Log(time.Now(), v)
		}
	})
}

func TestIsEmpty(t *testing.T) {
	tests.RunOnRedis(t, func(client redis.Client) {
		defer client.Close()

		store := newRedisStroe[delayTask](delay_queue, execute_queue, delay_task_set, client)
		store.Clear()

		assert.True(t, store.IsEmpty())

		store.Push(context.Background(), delayTask{Message: "bbb", At: time.Now().Add(time.Second)})
		assert.False(t, store.IsEmpty())
	}, redis.RedisOption{
		Addrs:    []string{"redis-10448.c90.us-east-1-3.ec2.cloud.redislabs.com:10448"},
		Password: "E7HFwvENEqimiB1EG4IjJSa2IUi0B22o",
	})
}
