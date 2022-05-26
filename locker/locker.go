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

func NewLocker() *sync.Mutex {
	return &sync.Mutex{}
}
