package locker

import "sync"

func NewRWLocker() *sync.RWMutex {
	return &sync.RWMutex{}
}
