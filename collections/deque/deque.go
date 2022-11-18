package deque

import "github.com/charlienet/go-mixed/locker"

type Deque[T any] struct {
	locker locker.RWLocker
}

func New[T any]() *Deque[T] {
	return &Deque[T]{
		locker: locker.EmptyLocker,
	}
}


