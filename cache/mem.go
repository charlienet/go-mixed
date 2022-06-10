package cache

import "time"

type MemCache interface {
	Get(key string) ([]byte, error)
	Set(key string, entry []byte, expire time.Duration) error
	Delete(key ...string) error
}
