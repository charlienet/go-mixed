package cache

import (
	"context"
	"time"

	"github.com/charlienet/go-mixed/cache/bigcache"
	"github.com/charlienet/go-mixed/cache/freecache"
	"github.com/charlienet/go-mixed/logx"
)

const defaultPrefix = "cache"


type option func(*Cache) error

type options struct {
	Prefix string
}

func WithRedis(opts RedisConfig) option {
	return func(c *Cache) error {
		if len(opts.Prefix) == 0 {
			opts.Prefix = defaultPrefix
		}

		rds := NewRedis(opts)

		c.rds = rds

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		return rds.Ping(ctx)
	}
}

func WithBigCache(opts bigcache.BigCacheConfig) option {
	return func(c *Cache) error {
		mem, err := bigcache.NewBigCache(opts)

		c.mem = mem
		return err
	}
}

func WithFreeCache(size int) option {
	return func(c *Cache) error {
		mem := freecache.NewFreeCache(size)
		c.mem = mem

		return nil
	}
}

// 使用自定义分布式缓存
func WithDistributedCache(rds DistributedCache) option {
	return func(c *Cache) error {
		c.rds = rds
		return nil
	}
}

func WithPublishSubscribe(p PublishSubscribe) option {
	return func(c *Cache) error {
		return nil
	}
}

func WithLogger(log logx.Logger) option {
	return func(c *Cache) error {
		c.logger = log

		return nil
	}
}

func acquireDefaultCache() *Cache {
	return &Cache{
		qps: NewQps(),
	}
}
