package redis

import (
	"context"
	_ "embed"
	"maps"
	"strings"
	"sync"
	"time"

	"github.com/charlienet/go-mixed/rand"
	"github.com/charlienet/go-mixed/redis"
	goredis "github.com/redis/go-redis/v9"
)

//go:embed redis_locker.lua
var redis_locker_function string

const (
	defaultExpire = time.Second * 20
	retryInterval = time.Millisecond * 10
)

var once sync.Once

type redis_locker_store struct {
	key     string
	sources map[string]string
	expire  time.Duration // 过期时间
	mu      sync.RWMutex
	clients []redis.Client
}

func NewRedisStore(key string, clients ...redis.Client) *redis_locker_store {
	once.Do(func() { redis.Clients(clients).LoadFunction(redis_locker_function) })

	locker := &redis_locker_store{
		key:     key,
		sources: make(map[string]string),
		clients: clients,
		expire:  defaultExpire,
	}

	go locker.expandLockTime()

	return locker
}

func (l *redis_locker_store) Lock(ctx context.Context, sourceName string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if l.TryLock(ctx, sourceName) {
				return nil
			}
		}

		time.Sleep(retryInterval)
	}
}

func (l *redis_locker_store) TryLock(ctx context.Context, sourceName string) bool {
	value := l.getSourceValue(sourceName)

	results := l.fCall(ctx, "locker_lock", sourceName, value, l.expire.Milliseconds())

	if !isSuccess(results) {
		for _, r := range results {
			if r.Err() != nil {
				println("err:", r.Err().Error())
			}
		}

		l.Unlock(ctx, sourceName)
		return false
	}

	return true
}

func (locker *redis_locker_store) Unlock(ctx context.Context, sourceName string) {
	value := locker.getSourceValue(sourceName)
	locker.fCall(ctx, "locker_unlock", sourceName, value)

	locker.mu.Lock()
	defer locker.mu.Unlock()

	delete(locker.sources, sourceName)
}

func (l *redis_locker_store) expandLockTime() {
	for {
		time.Sleep(l.expire / 3)

		if len(l.sources) == 0 {
			continue
		}

		l.mu.RLock()
		cloned := maps.Clone(l.sources)
		l.mu.RUnlock()

		for k, v := range cloned {
			results := l.fCall(context.Background(), "locker_expire", k, v, l.expire.Seconds())
			for _, r := range results {
				if r.Err() != nil {
					println("键延期失败:", r.Err().Error())
				}
			}
		}
	}
}

func (l *redis_locker_store) getSourceValue(name string) string {
	l.mu.Lock()
	defer l.mu.Unlock()

	if v, ok := l.sources[name]; ok {
		return v
	}

	v := rand.Hex.Generate(36)
	l.sources[name] = v
	return v
}

func (locker *redis_locker_store) fCall(ctx context.Context, cmd string, key string, args ...any) []*goredis.Cmd {
	results := make([]*goredis.Cmd, 0, len(locker.clients))

	var wg sync.WaitGroup
	wg.Add(len(locker.clients))
	for _, rdb := range locker.clients {
		go func(rdb redis.Client) {
			defer wg.Done()

			newKey := rdb.JoinKeys(locker.key, key)
			results = append(results, rdb.FCall(ctx, cmd, []string{newKey}, args...))
		}(rdb)
	}

	wg.Wait()

	return results
}

func isSuccess(results []*goredis.Cmd) bool {
	successCount := 0
	for _, ret := range results {
		resp, err := ret.Result()

		if err != nil || resp == nil {
			return false
		}

		reply, ok := resp.(string)
		if ok && strings.EqualFold(reply, "OK") {
			successCount++
		}
	}

	return successCount >= len(results)/2+1
}
