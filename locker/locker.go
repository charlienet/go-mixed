package locker

import "sync"

type Locker interface {
	Lock()
	Unlock()
	TryLock() bool
}

type RWLocker interface {
	Locker
	RLock()
	RUnlock()
	TryRLock() bool
}

type locker struct {
	*sync.Mutex
}

func NewLocker() *locker {
	return &locker{Mutex: &sync.Mutex{}}
}

