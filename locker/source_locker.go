package locker

import (
	"fmt"
)

// 资源锁
type SourceLocker struct {
	m     RWLocker
	locks map[string]Locker
}

func NewSourceLocker() *SourceLocker {
	return &SourceLocker{
		m:     NewRWLocker(),
		locks: make(map[string]Locker),
	}
}

func (s *SourceLocker) Lock(key string) {
	s.m.RLock()
	l, ok := s.locks[key]

	if ok {
		s.m.RUnlock()

		l.Lock()
		fmt.Println("加锁")
	} else {
		s.m.RUnlock()

		s.m.Lock()
		new := NewLocker()
		s.locks[key] = new
		s.m.Unlock()

		new.Lock()
		fmt.Println("初始加锁")
	}
}

func (s *SourceLocker) Unlock(key string) {
	s.m.Lock()
	if l, ok := s.locks[key]; ok {
		l.Unlock()
		// delete(s.locks, key)
		fmt.Println("解锁")
	}
	s.m.Unlock()
}

func (s *SourceLocker) TryLock(key string) bool {
	return false
}
