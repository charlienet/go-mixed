package cache

import (
	"time"

	"github.com/charlienet/go-mixed/locker"
	"github.com/vmihailenco/go-tinylfu"
)

var _ MemCache = &TinyLFU{}

type TinyLFU struct {
	mu  locker.Locker
	lfu *tinylfu.T
	ttl time.Duration
}

func NewTinyLFU(size int, ttl time.Duration) *TinyLFU {
	return &TinyLFU{
		mu:  locker.NewLocker(),
		lfu: tinylfu.New(size, 100000),
		ttl: ttl,
	}
}

func (c *TinyLFU) Set(key string, b []byte, expire time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lfu.Set(&tinylfu.Item{
		Key:      key,
		Value:    b,
		ExpireAt: time.Now().Add(c.ttl),
	})

	return nil
}

func (c *TinyLFU) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.lfu.Get(key)
	if !ok {
		return nil, false
	}

	return val.([]byte), true
}

func (c *TinyLFU) Delete(keys ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, k := range keys {
		c.lfu.Del(k)
	}

	return nil
}

func (c *TinyLFU) Clear() {

}
