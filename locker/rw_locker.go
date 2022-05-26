package locker

import "sync"

var _ RWLocker = &sync.RWMutex{}

func NewRWLocker() *sync.RWMutex {
	return &sync.RWMutex{}
}
