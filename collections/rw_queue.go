package collections

import "sync"

type rw_queue[T any] struct {
	q  Queue[T]
	mu sync.Mutex
}

func (q *rw_queue[T]) Push(v T) {
	q.mu.Lock()
	q.q.Put(v)
	q.mu.Unlock()
}

func (q *rw_queue[T]) Pop() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.q.Poll()
}

func (q *rw_queue[T]) Peek() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.q.Peek()
}

func (q *rw_queue[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.q.Size()
}

func (q *rw_queue[T]) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.q.IsEmpty()
}
