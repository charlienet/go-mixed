package cache

import (
	"errors"
	"time"

	"github.com/allegro/bigcache"
)

var _ MemCache = &bigCacheClient{}

type BigCacheConfig struct {
}

type bigCacheClient struct {
	cache *bigcache.BigCache
}

func NewBigCache(c *BigCacheConfig) (*bigCacheClient, error) {
	bigCache, err := bigcache.NewBigCache(bigcache.Config{})
	if err != nil {
		return nil, err
	}

	return &bigCacheClient{
		cache: bigCache,
	}, nil
}

func (c *bigCacheClient) Get(key string) ([]byte, error) {
	return c.cache.Get(key)
}

func (c *bigCacheClient) Set(key string, entry []byte, expire time.Duration) error {
	return c.cache.Set(key, entry)
}

func (c *bigCacheClient) Delete(key string) error {
	return c.cache.Delete(key)
}

func (c *bigCacheClient) Exist(key string) {
}

func (c *bigCacheClient) IsNotFound(err error) bool {
	return errors.Is(err, bigcache.ErrEntryNotFound)
}
