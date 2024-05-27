package locker

import "sync"

type locker interface {
	Lock()
	Unlock()
	TryLock() bool
}

type rwLocker interface {
	Lock()
	Unlock()
	TryLock() bool
	RLock()
	RUnlock()
	TryRLock() bool
}

func NewLocker() *sync.Mutex {
	return &sync.Mutex{}
}

func NewRWLocker() *sync.RWMutex {
	return &sync.RWMutex{}
}
