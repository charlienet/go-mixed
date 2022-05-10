package locker

import "sync"

var _ sync.Locker = &emptyLocker{}

type emptyLocker struct{}

func NewEmptyLocker() *emptyLocker {
	return &emptyLocker{}
}

func (l *emptyLocker) Lock() {}

func (l *emptyLocker) Unlock() {}
