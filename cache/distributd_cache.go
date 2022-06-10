package cache

import "time"

type DistributdCache interface {
	Get(key string, out any) error
	Set(key string, value any, expiration time.Duration) error
	Delete(key ...string) error
	Ping() error
}
