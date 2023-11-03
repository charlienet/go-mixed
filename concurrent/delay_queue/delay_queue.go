package delayqueue

import (
	"context"
	"time"

	"github.com/charlienet/go-mixed/locker"
)

type store[T Delayed] interface {
	Push(context.Context, T) error
	Pop() (T, error)
	Peek() (T, bool)
	IsEmpty() bool // 队列是否为空
}

type delayQueue[T Delayed] struct {
	mu    locker.RWLocker
	store store[T]
}

type Delayed interface {
	Delay() time.Time
}

func New[T Delayed]() *delayQueue[T] {
	return &delayQueue[T]{
		mu:    locker.NewRWLocker(),
		store: newMemStore[T](),
	}
}

func (q *delayQueue[T]) UseStore(s store[T]) *delayQueue[T] {
	q.store = s
	return q
}

func (q *delayQueue[T]) Push(task T) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.store.Push(context.Background(), task)
}

func (q *delayQueue[T]) Peek() (T, bool) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.store.Peek()
}

func (q *delayQueue[T]) Pop() (T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.store.Pop()
}

func (q *delayQueue[T]) Channel(size int) <-chan T {
	out := make(chan T, size)
	go func() {
		for {
			entry, _ := q.Pop()
			out <- entry
		}
	}()
	return out
}

func (q *delayQueue[T]) IsEmpty() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.store.IsEmpty()
}
