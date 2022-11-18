package locker

import (
	"fmt"
	"sync/atomic"
)

// 带计数器锁
type countLocker struct {
	Locker
	Count int32
}

// SourceLocker 资源锁
type SourceLocker struct {
	m     RWLocker
	locks map[string]*countLocker
}

func NewSourceLocker() *SourceLocker {
	return &SourceLocker{
		m:     NewRWLocker(),
		locks: make(map[string]*countLocker),
	}
}

func (s *SourceLocker) Lock(key string) {
	s.m.RLock()
	l, ok := s.locks[key]
	s.m.RUnlock()

	if ok {
		atomic.AddInt32(&l.Count, 1)
		l.Lock()

		fmt.Println("加锁")
	} else {
		// 加锁，再次检查是否已经具有锁
		s.m.Lock()
		if l2, ok := s.locks[key]; ok {
			s.m.Unlock()

			l2.Lock()
			fmt.Println("二次检查加锁")
		} else {
			n := NewLocker()
			s.locks[key] = &countLocker{Locker: n, Count: 1}

			s.m.Unlock()

			fmt.Printf("新锁准备加锁:%p\n", n)
			n.Lock()

			fmt.Println("初始加锁")
		}
	}
}

func (s *SourceLocker) Unlock(key string) {
	s.m.Lock()

	if l, ok := s.locks[key]; ok {
		atomic.AddInt32(&l.Count, -1)
		fmt.Printf("解锁%p\n", l)
		l.Unlock()

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

	if ok {
		ret := l.TryLock()
		s.m.RUnlock()

		return ret
	} else {
		s.m.RUnlock()

		s.m.Lock()
		n := NewLocker()
		s.locks[key] = &countLocker{Locker: n, Count: 1}
		s.m.Unlock()

		return n.TryLock()
	}
}
