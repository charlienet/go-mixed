package locker

import (
	"log"
	"sync"

	"github.com/charlienet/go-mixed/redis"
)

var empty = &emptyLocker{}

type Locker struct {
	once              sync.Once
	distributedLocker DistributedLocker // 分布式锁
	mu                locker
}

func (w *Locker) WithRedis(key string, rdb redis.Client) *Locker {
	return w
}

func (w *Locker) WithDistributedLocker(d DistributedLocker) *Locker {
	return w
}

func (w *Locker) Synchronize() *Locker {
	if w.mu == nil || w.mu == empty {
		w.mu = NewLocker()
	}

	return w
}

func (w *Locker) Lock() {
	w.ensureLocker().mu.Lock()
}

func (w *Locker) Unlock() {
	w.ensureLocker().mu.Unlock()
}

func (w *Locker) TryLock() bool {
	return w.ensureLocker().mu.TryLock()
}

func (w *Locker) ensureLocker() *Locker {
	w.once.Do(func() {
		if w.mu == nil {
			w.mu = empty
		}

	})

	return w
}

type SpinLocker struct {
	Locker
}

func (w *SpinLocker) Synchronize() {
	if w.mu == nil || w.mu == empty {
		w.mu = NewSpinLocker()
	}
}

type RWLocker struct {
	once sync.Once
	mu   rwLocker
}

func (w *RWLocker) Synchronize() *RWLocker {
	if w.mu == nil || w.mu == empty {
		w.mu = NewRWLocker()
	}

	return w
}

func (w *RWLocker) Lock() {
	w.ensureLocker().mu.Lock()
}

func (w *RWLocker) TryLock() bool {
	return w.ensureLocker().mu.TryLock()
}

func (w *RWLocker) Unlock() {
	w.ensureLocker().mu.Unlock()
}

func (w *RWLocker) RLock() {
	w.ensureLocker().mu.RLock()
}

func (w *RWLocker) TryRLock() bool {
	return w.ensureLocker().mu.TryRLock()
}

func (w *RWLocker) RUnlock() {
	w.ensureLocker().mu.RUnlock()
}

func (w *RWLocker) ensureLocker() *RWLocker {
	w.once.Do(func() {
		if w.mu == nil {
			log.Println("初始化一个空锁")
			w.mu = empty
		}
	})

	return w
}
