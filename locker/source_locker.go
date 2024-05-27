package locker

import (
	"context"
	"fmt"
	"sync/atomic"

	redis_store "github.com/charlienet/go-mixed/locker/redis"
	"github.com/charlienet/go-mixed/redis"
)

// 带计数器锁
type countLocker struct {
	rw    rwLocker
	Count int32
}

// SourceLocker 资源锁
type SourceLocker struct {
	m                 RWLocker
	distributedLocker DistributedLocker
	locks             map[string]*countLocker
	err               error
}

func NewSourceLocker() *SourceLocker {
	l := &SourceLocker{
		locks: make(map[string]*countLocker),
	}

	l.m.Synchronize()
	return l
}

func (s *SourceLocker) WithRedis(key string, clients ...redis.Client) *SourceLocker {
	redisStore := redis_store.NewRedisStore(key, clients...)
	return s.WithDistributedLocker(redisStore)
}

func (s *SourceLocker) WithDistributedLocker(distributed DistributedLocker) *SourceLocker {
	s.distributedLocker = distributed
	return s
}

func (s *SourceLocker) Lock(key string) {
	s.m.RLock()
	l, ok := s.locks[key]
	s.m.RUnlock()

	if ok {
		atomic.AddInt32(&l.Count, 1)
		l.rw.Lock()

		fmt.Println("加锁")
	} else {
		// 加锁，再次检查是否已经具有锁
		s.m.Lock()
		if l2, ok := s.locks[key]; ok {
			s.m.Unlock()

			l2.rw.Lock()
			fmt.Println("二次检查加锁")
		} else {
			n := NewRWLocker()
			s.locks[key] = &countLocker{rw: n, Count: 1}

			s.m.Unlock()

			n.Lock()
		}
	}
}

func (s *SourceLocker) Unlock(key string) {
	s.m.Lock()

	if l, ok := s.locks[key]; ok {
		atomic.AddInt32(&l.Count, -1)
		l.rw.Unlock()

		if s.distributedLocker != nil {
			s.distributedLocker.Unlock(context.Background(), key)
		}

		if l.Count == 0 {
			delete(s.locks, key)
		}
	}
	s.m.Unlock()
}

func (s *SourceLocker) TryLock(key string) bool {
	// 加读锁
	s.m.RLock()
	l, ok := s.locks[key]
	s.m.RUnlock()
	if ok {
		ret := l.rw.TryLock()
		return ret
	} else {
		s.m.Lock()
		n := NewRWLocker()
		s.locks[key] = &countLocker{rw: n, Count: 1}
		s.m.Unlock()

		return n.TryLock()
	}
}
