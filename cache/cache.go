package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/charlienet/go-mixed/json"
	"github.com/charlienet/go-mixed/locker"
	"github.com/charlienet/go-mixed/logx"
)

var ErrNotFound = errors.New("key not found")

type LoadFunc func(context.Context) (any, error)

type Cache struct {
	prefix           string            // 键前缀
	retry            int               // 资源获取时的重试次数
	mem              MemCache          // 内存缓存
	distributdCache  DistributdCache   // 分布式缓存
	publishSubscribe PublishSubscribe  // 发布订阅
	lock             locker.ChanLocker // 资源锁
	stats            *Stats            // 缓存命中计数
	qps              *qps              // 访问计数
	logger           logx.Logger       // 日志记录
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
	// 从缓存加载数据
	// 1. 检查内存是否存在
	// 2. 检查分布缓存是否存在
}

// 从数据源加载数据
func (c *Cache) getFromSource(ctx context.Context, key string, fn LoadFunc) error {

	// 1. 尝试获取资源锁，如成功获取到锁加载数据
	// 2. 未获取到锁，等待从缓存中获取
	ch, ok := c.lock.Get(key)
	if ok {
		defer c.lock.Release(key)

		v, err := fn(ctx)
		if err != nil {
			return fmt.Errorf("load from source err:%v", err)
		}

		// 取出值存入多级缓存
		_ = v

		return nil
	}

	// 等待数据加载完成
	select {
	case <-ch:

		// 未取到结果时，再次获取
		return c.getFromSource(ctx, key, fn)
	}
}

func (c *Cache) marshal(value any) ([]byte, error) {
	switch value := value.(type) {
	case nil:
		return nil, nil
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	}

	b, err := json.Marshal(value)
	return b, err
}

func (c *Cache) unmarshal(b []byte, value any) error {
	if len(b) == 0 {
		return nil
	}

	switch value := value.(type) {
	case nil:
		return nil
	case *[]byte:
		clone := make([]byte, len(b))
		copy(clone, b)
		*value = clone
		return nil
	case *string:
		*value = string(b)
		return nil
	}

	err := json.Unmarshal(b, value)
	return err
}
