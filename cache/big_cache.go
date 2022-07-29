package cache

import (
	"errors"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/charlienet/go-mixed/logx"
)

var _ MemCache = &bigCacheClient{}

type BigCacheConfig struct {
	Shards             int
	LifeWindow         time.Duration
	CleanWindow        time.Duration
	MaxEntriesInWindow int
	MaxEntrySize       int
	HardMaxCacheSize   int
	log                logx.Logger
}

type bigCacheClient struct {
	cache *bigcache.BigCache
}

func NewBigCache(c BigCacheConfig) (*bigCacheClient, error) {
	config := bigcache.DefaultConfig(time.Minute * 10)

	config.LifeWindow = c.LifeWindow
	config.LifeWindow = c.LifeWindow
	config.CleanWindow = c.CleanWindow
	config.MaxEntriesInWindow = c.MaxEntriesInWindow
	config.MaxEntrySize = c.MaxEntrySize
	config.HardMaxCacheSize = c.HardMaxCacheSize
	config.Logger = c.log

	if c.Shards > 0 {
		config.Shards = c.Shards
	}

	bigCache, err := bigcache.NewBigCache(config)
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

func (c *bigCacheClient) Delete(keys ...string) error {
	ks := keys[:]
	for _, k := range ks {
		if err := c.cache.Delete(k); err != nil {
			return err
		}
	}

	return nil
}

func (c *bigCacheClient) Exist(key string) {
}

func (c *bigCacheClient) IsNotFound(err error) bool {
	return errors.Is(err, bigcache.ErrEntryNotFound)
}
