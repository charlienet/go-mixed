package freecache

import (
	"errors"
	"time"

	"github.com/charlienet/go-mixed/bytesconv"
	"github.com/coocood/freecache"
)

const defaultSize = 10 * 1024 * 1024 // 10M

type freeCache struct {
	cache *freecache.Cache
}

func NewFreeCache(size int) *freeCache {
	if size < defaultSize {
		size = defaultSize
	}

	// debug.SetGCPercent(20)

	c := freecache.NewCache(size)
	return &freeCache{
		cache: c,
	}
}

func (c *freeCache) Get(key string) ([]byte, bool) {
	b, err := c.cache.Get([]byte(key))
	if err != nil {
		return b, false
	}
	return b, true
}

func (c *freeCache) Set(key string, value []byte, d time.Duration) error {
	s := int(d.Seconds())
	return c.cache.Set([]byte(key), value, s)
}

func (c *freeCache) Delete(keys ...string) error {
	for _, k := range keys {
		affected := c.cache.Del(bytesconv.StringToBytes(k))

		if !affected {
			return errors.New("不存在")
		}
	}

	return nil
}

func (c *freeCache) Exist(key string) error {
	return nil
}

func (c *freeCache) Clear() {

}

func (c *freeCache) IsNotFound(err error) bool {
	return errors.Is(err, freecache.ErrNotFound)
}
