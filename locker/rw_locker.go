package locker

import "sync"

var _ RWLocker = &rwLocker{}

type rwLocker struct {
	*sync.RWMutex
}

func NewRWLocker() *rwLocker {
	return &rwLocker{RWMutex: &sync.RWMutex{}}
}
