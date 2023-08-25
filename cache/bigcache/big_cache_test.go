package bigcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBigCache(t *testing.T) {
	r := require.New(t)

	c, err := NewBigCache(BigCacheConfig{})
	r.Nil(err)

	cacheKey := "a"
	cacheValue := "bbb"

	c.Set(cacheKey, []byte(cacheValue), time.Second*5)

	r.True(c.Exist(cacheKey))
	r.False(c.Exist("abb"))

	b, ok := c.Get(cacheKey)
	r.True(ok)
	r.Equal(cacheValue, string(b))
}
