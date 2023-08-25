package cache

import "context"

// var emptyCache DistributedCache = &emptyCacheAdapter{}

type emptyCacheAdapter struct {
}

func (*emptyCacheAdapter) Delete(ctx context.Context, keys ...string) {}
