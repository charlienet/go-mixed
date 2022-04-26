package cache

import (
	"errors"
	"time"

	"github.com/charlienet/go-mixed/bytesconv"
)

var ErrNotFound = errors.New("not found")

type LoadFunc func() (any, error)

type Cache struct {
	prefix           string           // 键前缀
	mem              MemCache         // 内存缓存
	distributdCache  DistributdCache  // 分布式缓存
	publishSubscribe PublishSubscribe // 发布订阅
	qps              *qps
}

func NewCache(opts ...option) (*Cache, error) {
	c := &Cache{
		qps: NewQps(),
	}

	for _, f := range opts {
		f(c)
	}

	go c.subscribe()

	return c, nil
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

func (c *Cache) GetFn(key string, out any, fn LoadFunc, expiration time.Duration) (bool, error) {
	ret, err := fn()
	if err != nil {
		return false, err
	}

	_ = ret

	return false, nil
}

func (c *Cache) Exist(key string) (bool, error) {
	return false, nil
}

func (c *Cache) Delete(key string) error {
	if c.mem != nil {
		c.mem.Delete(key)
	}

	if c.distributdCache != nil {
		c.distributdCache.Delete(key)
	}

	return nil
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

func (c *Cache) subscribe() {
}

func (c *Cache) genKey(key string) string {
	if len(c.prefix) == 0 {
		return key
	}

	return c.prefix + "-" + key
}
