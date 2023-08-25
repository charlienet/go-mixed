package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/charlienet/go-mixed/locker"
	"github.com/charlienet/go-mixed/logx"
	"golang.org/x/sync/singleflight"
)

var ErrNotFound = errors.New("key not found")

// 数据加载函数定义
type LoadFunc func(context.Context) (any, error)

type ICache interface {
}

type Cache struct {
	prefix           string             // 键前缀
	retry            int                // 资源获取时的重试次数
	mem              MemCache           // 内存缓存
	rds              DistributedCache   // 远程缓存
	publishSubscribe PublishSubscribe   // 发布订阅
	group            singleflight.Group // singleflight.Group
	lock             locker.ChanLocker  // 资源锁
	stats            *Stats             // 缓存命中计数
	qps              *qps               // 访问计数
	logger           logx.Logger        // 日志记录
}

func New(opts ...option) (*Cache, error) {

	c := acquireDefaultCache()
	for _, f := range opts {
		if err := f(c); err != nil {
			return c, nil
		}
	}

	// 未设置内存缓存时，添加默认缓存
	if c.mem == nil {
		c.mem = NewTinyLFU(1<<12, time.Second*30)
	}

	return c, nil
}

func (c *Cache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	buf, err := Marshal(value)
	if err != nil {
		return err
	}

	if c.mem != nil {
		c.mem.Set(key, buf, expiration)
	}

	if c.rds != nil {
		c.rds.Set(ctx, key, buf, expiration)
	}

	return nil
}

func (c *Cache) Get(ctx context.Context, key string, out any) error {
	if c.mem != nil {
		c.getFromMem(key)
	}

	if c.rds != nil {
		if err := c.rds.Get(ctx, key, out); err != nil {

		}
	}

	return nil
}

func (c *Cache) GetFn(ctx context.Context, key string, out any, fn LoadFunc, expiration time.Duration) (bool, error) {
	c.Get(ctx, key, out)

	// 多级缓存中未找到时,放置缓存对象
	ret, err := fn(ctx)
	if err != nil {
		return false, err
	}

	c.Set(ctx, key, ret, expiration)

	return false, nil
}

func (c *Cache) Exist(key string) (bool, error) {
	return false, nil
}

func (c *Cache) Delete(ctx context.Context, key ...string) error {
	if c.mem != nil {
		c.mem.Delete(key...)
	}

	if c.rds != nil {
		c.rds.Delete(ctx, key...)
	}

	for _, k := range key {
		c.group.Forget(k)
	}

	return nil
}

// 清除本地缓存
func (c *Cache) ClearMem() {
	if c.mem != nil {
		c.mem.Clear()
	}
}

func (c *Cache) Clear() {

}

func (c *Cache) Disable() {

}

func (c *Cache) Enable() {

}

func (c *Cache) getOnce(ctx context.Context, key string) (b []byte, cached bool, err error) {
	if c.mem != nil {
		b, ok := c.mem.Get(key)
		if ok {
			return b, true, nil
		}
	}
	c.group.Do(key, func() (any, error) {
		if c.mem != nil {
			b, ok := c.mem.Get(key)
			if ok {
				return b, nil
			}
		}

		if c.rds != nil {
			c.rds.Get(ctx, key, nil)
		}

		return nil, nil
	})

	return
}

func (c *Cache) getFromMem(key string) ([]byte, bool) {
	bytes, cached := c.mem.Get(key)
	return bytes, cached
}

// 从缓存加载数据
func (c *Cache) getFromCache() {
	// 从缓存加载数据
	// 1. 检查内存是否存在
	// 2. 检查分布缓存是否存在
}

// 从数据源加载数据
func (c *Cache) getFromSource(ctx context.Context, key string, fn LoadFunc) error {

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
