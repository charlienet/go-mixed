package cache

type option func(*Cache)

type options struct {
	Prefix string
}

func WithPrefix(prefix string) option {
	return func(o *Cache) { o.prefix = prefix }
}

func WithDistributdCache(d DistributdCache) option {
	return func(o *Cache) { o.distributdCache = d }
}

func WithBigCache(config *BigCacheConfig) option {
	return func(o *Cache) {
		c, err := NewBigCache(config)
		_ = err
		o.mem = c
	}
}

func WithFreeCache(size int) option {
	return func(o *Cache) { o.mem = NewFreeCache(size) }
}

func WithPublishSubscribe(p PublishSubscribe) option {
	return func(o *Cache) {}
}
