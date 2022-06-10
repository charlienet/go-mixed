package cache

import (
	"context"
	"errors"
	"time"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/logx"
)

var ErrNotFound = errors.New("key not found")

type LoadFunc func(context.Context) (any, error)

type Cache struct {
	prefix           string           // 键前缀
	retry            int              // 资源获取时的重试次数
	mem              MemCache         // 内存缓存
	distributdCache  DistributdCache  // 分布式缓存
	publishSubscribe PublishSubscribe // 发布订阅
	qps              *qps             //
	logger           logx.Logger      // 日志记录
}

func NewCache(opts ...option) *Cache {

	c := acquireDefaultCache()
	for _, f := range opts {
		if err := f(c); err != nil {
			return c
		}
	}

	go c.subscribe()

	return c
}

func (c *Cache) Set(key string, value any, expiration time.Duration) error {
	if c.mem != nil {
		bytes, err := bytesconv.Encode(value)
		if err != nil {
			return err
		}

		c.mem.Set(key, bytes, expiration)
	}

	return nil
}

func (c *Cache) Get(key string, out any) error {
	if c.mem != nil {
		c.getFromMem(key, out)
	}

	if c.distributdCache != nil {
		if err := c.distributdCache.Get(key, out); err != nil {

		}
	}

	return nil
}

func (c *Cache) GetFn(ctx context.Context, key string, out any, fn LoadFunc, expiration time.Duration) (bool, error) {
	c.Get(key, out)

	// 多级缓存中未找到时,放置缓存对象
	ret, err := fn(ctx)
	if err != nil {
		return false, err
	}

	c.Set(key, ret, expiration)

	return false, nil
}

func (c *Cache) Exist(key string) (bool, error) {
	return false, nil
}

func (c *Cache) Delete(key ...string) error {
	if c.mem != nil {
		c.mem.Delete(key...)
	}

	if c.distributdCache != nil {
		c.distributdCache.Delete(key...)
	}

	return nil
}

func (c *Cache) subscribe() {
}

func (c *Cache) getFromMem(key string, out any) error {
	bytes, err := c.mem.Get(key)
	if err != nil {
		return err
	}

	if err := bytesconv.Decode(bytes, out); err != nil {
		return err
	}

	return nil
}

// 从缓存加载数据
func (c *Cache) getFromCache() {

}

// 从数据源加载数据
func (c *Cache) getFromSource(ctx context.Context, key string, fn LoadFunc) {

	// 1. 尝试获取资源锁，如成功获取到锁加载数据
	// 2. 未获取到锁，等待从缓存中获取
	fn(ctx)

}
