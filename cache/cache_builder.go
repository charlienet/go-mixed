package cache

import "github.com/charlienet/go-mixed/logx"

const defaultPrefix = "cache"

type option func(*Cache) error

type options struct {
	Prefix string
}

func acquireDefaultCache() *Cache {
	return &Cache{
		prefix: defaultPrefix,
		qps:    NewQps(),
	}
}

type cacheBuilder struct {
	prefix           string
	redisOptions     RedisConfig
	bigCacheConfig   BigCacheConfig
	freeSize         int
	publishSubscribe PublishSubscribe
	log              logx.Logger
}

func NewCacheBuilder() *cacheBuilder {
	return &cacheBuilder{}
}

func (b *cacheBuilder) WithLogger(log logx.Logger) *cacheBuilder {
	b.log = log
	return b
}

func (b *cacheBuilder) WithPrefix(prefix string) *cacheBuilder {
	b.prefix = prefix
	return b
}

func (b *cacheBuilder) WithRedis(opts RedisConfig) *cacheBuilder {
	b.redisOptions = opts
	return b
}

func (b *cacheBuilder) WithBigCache(opts BigCacheConfig) *cacheBuilder {
	b.bigCacheConfig = opts
	return b
}

func (b *cacheBuilder) WithFreeCache(size int) *cacheBuilder {
	b.freeSize = size
	return b
}

// 使用自定义分布式缓存
func WithDistributedCache(c DistributdCache) {

}

func (b *cacheBuilder) WithPublishSubscribe(p PublishSubscribe) *cacheBuilder {
	b.publishSubscribe = p
	return b
}

func (b cacheBuilder) Build() (*Cache, error) {
	var err error
	cache := acquireDefaultCache()
	if len(b.prefix) > 0 {
		cache.prefix = b.prefix
	}

	b.redisOptions.Prefix = cache.prefix

	redis := NewRedis(b.redisOptions)
	if err := redis.Ping(); err != nil {
		return cache, err
	}

	var mem MemCache
	if b.freeSize > 0 {
		mem = NewFreeCache(b.freeSize)
	} else {
		if b.log != nil {
			b.bigCacheConfig.log = b.log
		}

		mem, err = NewBigCache(b.bigCacheConfig)
	}

	cache.distributdCache = redis
	cache.mem = mem
	cache.publishSubscribe = b.publishSubscribe
	cache.logger = b.log

	return cache, err
}
