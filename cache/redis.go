package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/json"
	"github.com/charlienet/go-mixed/rand"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Perfix string // key perfix
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
	perfix     string // 缓存键前缀
}

func NewRedis(c *RedisConfig) *redisClient {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    c.Addrs,
		DB:       c.DB,
		Username: c.Username,
		Password: c.Password,
	})

	return &redisClient{
		emptyStamp: fmt.Sprintf("redis-empty-%d-%s", time.Now().Unix(), rand.Hex.Generate(6)),
		perfix:     c.Perfix,
		client:     client,
	}
}

func (c *redisClient) Get(key string, out any) error {
	cmd := c.client.Get(context.Background(), key)
	str, err := cmd.Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytesconv.StringToBytes(str), out)
	return err
}

func (c *redisClient) Set(key string, value any, expiration time.Duration) {
	j, _ := json.Marshal(value)
	c.client.Set(context.Background(), key, j, expiration)
}

func (c *redisClient) Exist(key string) (bool, error) {
	return false, nil
}

func (c *redisClient) Delete(key string) error {
	cmd := c.client.Del(context.Background(), key)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (c *redisClient) Ping() error {
	_, err := c.client.Ping(context.Background()).Result()
	return err
}
