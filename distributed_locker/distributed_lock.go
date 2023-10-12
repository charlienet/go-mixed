package locker

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/charlienet/go-mixed/rand"
	"github.com/charlienet/go-mixed/redis"
	goredis "github.com/redis/go-redis/v9"
)

const (
	// 加锁(可重入)
	lockCmd = `if redis.call("GET", KEYS[1]) == ARGV[1] then
	redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
	return "OK"
else
	return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`

	// 解锁
	delCmd = `if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return '0'
end`

	// 延期
	incrCmd = `
if redis.call('get', KEYS[1]) == ARGV[1] then
  return redis.call('expire', KEYS[1], ARGV[2])
 else
   return '0'
end`
)

const (
	defaultExpire = time.Second * 10
	retryInterval = time.Millisecond * 10
)

var (
	once             sync.Once
	ErrContextCancel = errors.New("context cancel")
)

type distributedlock struct {
	clients  []redis.Client  // redis 客户端
	ctx      context.Context //
	key      string          // 资源键
	rand     string          // 随机值
	unlocked bool            // 是否已解锁
	expire   time.Duration   // 过期时间
}

func NewDistributedLocker(ctx context.Context, key string, clients ...redis.Client) *distributedlock {
	expire := defaultExpire
	if deadline, ok := ctx.Deadline(); ok {
		expire = deadline.Sub(time.Now())
	}

	locker := &distributedlock{
		ctx:     ctx,
		clients: clients,
		key:     key,
		rand:    rand.Hex.Generate(24),
		expire:  expire,
	}

	return locker
}

func (locker *distributedlock) Lock() error {
	for {
		select {
		case <-locker.ctx.Done():
			return ErrContextCancel
		default:
			if locker.TryLock() {
				return nil
			}
		}

		time.Sleep(retryInterval)
	}
}

func (locker *distributedlock) TryLock() bool {
	results := locker.Eval(locker.ctx, lockCmd, []string{locker.key}, locker.rand, locker.expire.Milliseconds())

	if !isSuccess(results) {
		locker.Unlock()
		return false
	}

	locker.expandLockTime()

	return true
}

func (locker *distributedlock) Unlock() {
	locker.Eval(locker.ctx, delCmd, []string{locker.key}, locker.rand)
	locker.unlocked = true
}

func (l *distributedlock) expandLockTime() {
	once.Do(func() {
		go func() {
			for {
				time.Sleep(l.expire / 3)
				if l.unlocked {
					break
				}

				l.resetExpire()
			}
		}()
	})
}

func (locker *distributedlock) resetExpire() {
	locker.Eval(locker.ctx,
		incrCmd,
		[]string{locker.key},
		locker.rand,
		locker.expire.Seconds())
}

func (locker *distributedlock) Eval(ctx context.Context, cmd string, keys []string, args ...any) []*goredis.Cmd {
	results := make([]*goredis.Cmd, 0, len(locker.clients))

	var wg sync.WaitGroup
	wg.Add(len(locker.clients))
	for _, rdb := range locker.clients {
		go func(rdb redis.Client) {
			defer wg.Done()

			results = append(results, rdb.Eval(ctx, cmd, keys, args...))
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
