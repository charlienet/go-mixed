package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	defaultSeparator = ":"

	blockingQueryTimeout = 5 * time.Second
	readWriteTimeout     = 2 * time.Second
	defaultSlowThreshold = time.Millisecond * 100 // 慢查询
)

type Option func(r *Redis)

type Redis struct {
	addr      string // 服务器地址
	prefix    string // 键值前缀
	separator string // 分隔符
}

func New(addr string, opts ...Option) *Redis {
	r := &Redis{
		addr: addr,
	}

	return r
}

func (s *Redis) Set(ctx context.Context, key, value string) error {
	conn, err := s.getRedis()
	if err != nil {
		return err
	}

	return conn.Set(ctx, s.formatKey(key), value, 0).Err()
}

func (s *Redis) Get(ctx context.Context, key string) (string, error) {
	conn, err := s.getRedis()
	if err != nil {
		return "", err
	}

	return conn.Get(ctx, s.formatKey(key)).Result()
}

func (s *Redis) GetSet(ctx context.Context, key, value string) (string, error) {
	conn, err := s.getRedis()
	if err != nil {
		return "", err
	}

	val, err := conn.GetSet(ctx, s.formatKey(key), value).Result()
	return val, err
}

func (s *Redis) Del(ctx context.Context, key ...string) (int, error) {
	conn, err := s.getRedis()
	if err != nil {
		return 0, err
	}

	keys := s.formatKeys(key...)
	v, err := conn.Del(ctx, keys...).Result()
	if err != nil {
		return 0, err
	}

	return int(v), err
}

func (s *Redis) getRedis() (redis.UniversalClient, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{s.addr},
	})

	return client, nil
}

func (s *Redis) formatKeys(keys ...string) []string {
	// If no prefix is configured, this parameter is returned
	if s.prefix == "" {
		return keys
	}

	ret := make([]string, 0, len(keys))
	for _, k := range keys {
		ret = append(ret, s.formatKey(k))
	}

	return ret
}

func (s *Redis) formatKey(key string) string {
	if s.prefix == "" {
		return key
	}

	return s.prefix + s.separator + key
}
