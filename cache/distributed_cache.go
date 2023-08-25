package cache

import (
	"context"
	"time"
)

// 分布式缓存接口
type DistributedCache interface {
	Get(ctx context.Context, key string, out any) error
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Delete(ctx context.Context, key ...string) error
	Ping(ctx context.Context) error
}
