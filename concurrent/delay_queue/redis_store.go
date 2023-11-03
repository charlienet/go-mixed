package delayqueue

import (
	"context"
	"encoding"
	"encoding/json"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/charlienet/go-mixed/hash"
	"github.com/charlienet/go-mixed/redis"
)

// 使用Redis存储队列

type redisStore[T Delayed] struct {
	rdb          redis.Client
	delayQueue   string
	executeQueue string
	delayTaskSet string
}

func (q *delayQueue[T]) UseRedis(delayQueueName, executeQueueName, delayTaskName string, rdb redis.Client) *delayQueue[T] {
	q.store = newRedisStroe[T](delayQueueName, executeQueueName, delayTaskName, rdb)
	return q
}
func newRedisStroe[T Delayed](delayQueueName, executeQueueName, delayTaskName string, rdb redis.Client) *redisStore[T] {

	store := &redisStore[T]{
		delayQueue:   delayQueueName,
		executeQueue: executeQueueName,
		delayTaskSet: delayTaskName,

		rdb: rdb,
	}

	go func() {
		for {
			store.pushToExecute()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	return store
}

func (s *redisStore[T]) Push(ctx context.Context, v T) error {
	o := any(v).(encoding.BinaryMarshaler)
	bytes, err := o.MarshalBinary()
	if err != nil {
		return err
	}

	tx := s.rdb.TxPipeline()
	tx.HSet(context.Background(), s.delayTaskSet, hash.Sha1(bytes).Hex(), bytes)
	tx.Exec(context.Background())

	tx.HSet(context.Background(), s.delayTaskSet)

	// tx.Exec()
	ret := s.rdb.ZAdd(ctx, s.delayQueue, goredis.Z{
		Score:  float64(v.Delay().Unix()),
		Member: v,
	})

	return ret.Err()
}

func (s *redisStore[T]) pushToExecute() error {
	now := time.Now().Unix()

	ret, err := s.rdb.ZRangeByScore(
		context.Background(),
		s.delayQueue,
		&goredis.ZRangeBy{
			Min: "-inf",
			Max: strconv.FormatInt(now, 10),
		}).Result()

	if err != nil {
		return err
	}

	if len(ret) > 0 {
		pipe := s.rdb.TxPipeline()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		pipe.LPush(ctx, s.executeQueue, ret)
		pipe.ZRem(ctx, s.delayQueue, ret)

		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s *redisStore[T]) Pop() (T, error) {
	for {
		v, err := s.rdb.RPop(context.Background(), s.executeQueue).Result()
		if err != nil {
			if err == redis.Nil {
				time.Sleep(time.Millisecond * 10)
				continue
			}

			return *new(T), err
		}

		if len(v) > 0 {
			var task T
			if err := json.Unmarshal([]byte(v), &task); err != nil {
				return *new(T), err
			}

			return task, nil
		}
	}
}

func (s *redisStore[T]) Peek() (t T, r bool) {
	m, err := s.rdb.ZRange(context.Background(), s.delayQueue, 0, 0).Result()
	if err != nil {
		return *new(T), false
	}

	if len(m) == 1 {
		var t T

		s := m[0]
		if err := json.Unmarshal([]byte(s), &t); err != nil {
			return *new(T), false
		}
		return t, true
	}

	return *new(T), false
}

func (s *redisStore[T]) Clear() {
	s.rdb.Del(context.Background(), s.delayQueue)
	s.rdb.Del(context.Background(), s.executeQueue)
}

func (s *redisStore[T]) IsEmpty() bool {
	n, _ := s.rdb.LLen(context.Background(), s.executeQueue).Result()
	m, _ := s.rdb.ZCard(context.Background(), s.delayQueue).Result()

	return (m + n) == 0
}
