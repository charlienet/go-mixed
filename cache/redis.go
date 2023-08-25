package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/json"
	"github.com/charlienet/go-mixed/rand"
	"github.com/go-redis/redis/v8"
)

const redisEmptyObject = "redis object not exist"

type RedisConfig struct {
	Prefix string // key perfix
	Addrs  []string

	// Database to be selected after connecting to the server.
	// Only single-node and failover clients.
	DB int

	Username        string
	Password        string
	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type redisClient struct {
	client     redis.UniversalClient
	emptyStamp string // 空对象标识，每个实例隔离
	prefix     string // 缓存键前缀
}

func NewRedis(c RedisConfig) *redisClient {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    c.Addrs,
		DB:       c.DB,
		Username: c.Username,
		Password: c.Password,
	})

	return &redisClient{
		emptyStamp: fmt.Sprintf("redis-empty-%d-%s", time.Now().Unix(), rand.Hex.Generate(6)),
		prefix:     c.Prefix,
		client:     client,
	}
}

func (c *redisClient) Get(cxt context.Context, key string, out any) error {
	val, err := c.client.Get(context.Background(), c.getKey(key)).Result()
	if errors.Is(err, redis.Nil) {
		return ErrNotFound
	}

	if err != nil {
		return err
	}

	// redis 保存键为空值时返回键不存在错误
	if val == redisEmptyObject {
		return ErrNotFound
	}

	return json.Unmarshal(bytesconv.StringToBytes(val), out)
}

func (c *redisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	j, _ := json.Marshal(value)
	return c.client.Set(context.Background(), c.getKey(key), j, expiration).Err()
}

func (c *redisClient) Exist(ctx context.Context, key string) (bool, error) {
	val, err := c.client.Exists(context.Background(), c.getKey(key)).Result()
	return val > 0, err
}

func (c *redisClient) Delete(ctx context.Context, key ...string) error {
	keys := make([]string, 0, len(key))
	for _, k := range key {
		keys = append(keys, c.getKey(k))
	}

	_, err := c.client.Del(context.Background(), keys...).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *redisClient) Ping(ctx context.Context) error {
	_, err := c.client.Ping(ctx).Result()
	return err
}

func (c *redisClient) getKey(key string) string {
	if c.prefix != "" {
		return c.prefix + ":" + key
	}

	return key
}
