package cache

import "time"

type MemCache interface {
	Get(key string) ([]byte, bool)
	Set(key string, b []byte, expire time.Duration) error
	Delete(key ...string) error
	Clear()
}
