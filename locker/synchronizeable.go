package locker

import (
	"log"
	"sync"
)

type WithLocker struct {
	once sync.Once
	mu   Locker
}

func (w *WithLocker) Synchronize() {
	if w.mu == nil || w.mu == EmptyLocker {
		w.mu = NewLocker()
	}
}

func (w *WithLocker) Lock() {
	w.ensureLocker().Lock()
}

func (w *WithLocker) Unlock() {
	w.ensureLocker().Unlock()
}

func (w *WithLocker) TryLock() bool {
	return w.ensureLocker().TryLock()
}

func (w *WithLocker) ensureLocker() Locker {
	w.once.Do(func() {
		if w.mu == nil {
			w.mu = EmptyLocker
		}
	})

	return w.mu
}

type WithSpinLocker struct {
	WithLocker
}

func (w *WithSpinLocker) Synchronize() {
	if w.mu == nil || w.mu == EmptyLocker {
		w.mu = NewSpinLocker()
	}
}

type WithRWLocker struct {
	once sync.Once
	mu   RWLocker
}

func (w *WithRWLocker) Synchronize() {
	if w.mu == nil || w.mu == EmptyLocker {
		log.Println("初始化有效锁")
		w.mu = NewRWLocker()
	}
}

func (w *WithRWLocker) Lock() {
	w.ensureLocker().Lock()
}

func (w *WithRWLocker) TryLock() bool {
	return w.ensureLocker().TryLock()
}

func (w *WithRWLocker) Unlock() {
	w.ensureLocker().Unlock()
}

func (w *WithRWLocker) RLock() {
	w.ensureLocker().RLock()
}

func (w *WithRWLocker) TryRLock() bool {
	return w.ensureLocker().TryRLock()
}

func (w *WithRWLocker) RUnlock() {
	w.ensureLocker().RUnlock()
}

func (w *WithRWLocker) ensureLocker() RWLocker {
	w.once.Do(func() {
		if w.mu == nil {
			log.Println("初始化一个空锁")
			w.mu = EmptyLocker
		}
	})

	return w.mu
}
