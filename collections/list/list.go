package list

import (
	"errors"

	"github.com/charlienet/go-mixed/locker"
)

var ErrorOutOffRange = errors.New("out of range")

type List[T any] interface {
}

type list[T any] struct {
	size   int
	locker locker.RWLocker
}

func (l *list[T]) Synchronize() {
	l.locker = locker.NewRWLocker()
}

func (l *list[T]) ForEach(fn func(T) bool) { panic("Not Implemented") }

func (l *LinkedList[T]) ToSlice() []T {
	s := make([]T, 0, l.Size())
	l.ForEach(func(t T) bool {
		s = append(s, t)
		return false
	})

	return s
}

func (l *list[T]) Size() int { return l.size }
