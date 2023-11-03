package delayqueue

import (
	"container/heap"
	"context"
	"sync"
	"time"
)

type delayedQueue []Delayed

type memStore[T Delayed] struct {
	mu     sync.Mutex
	h      *delayedQueue
	wakeup chan struct{}
}

func (q delayedQueue) Len() int {
	return len(q)
}

func (q delayedQueue) Less(i, j int) bool {
	return q[i].Delay().Before(q[j].Delay())
}

func (q delayedQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *delayedQueue) Push(x any) {
	*q = append(*q, x.(Delayed))
}

func (h *delayedQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func newMemStore[T Delayed]() *memStore[T] {
	store := &memStore[T]{
		h:      new(delayedQueue),
		wakeup: make(chan struct{}, 1),
	}

	heap.Init(store.h)
	return store
}

func (s *memStore[T]) Push(ctx context.Context, v T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	heap.Push(s.h, v)
	return nil
}

func (s *memStore[T]) Pop() (T, error) {
	for {
		_, exist := s.Peek()
		if exist {
			return s.h.Pop().(T), nil
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func (s *memStore[T]) Peek() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.h.Len() > 0 {
		return (*s.h)[0].(T), true
	}

	return *new(T), false
}

func (s *memStore[T]) Len() int {
	return s.h.Len()
}

func (s *memStore[T]) IsEmpty() bool {
	return s.Len() == 0
}
