package queue

import (
	"github.com/charlienet/go-mixed/collections/list"
)

type Queue[T any] struct {
	list list.LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Synchronize() *Queue[T] {
	q.list.Synchronize()

	return q
}

func (q *Queue[T]) Push(v T) {

}

func (q *Queue[T]) Pop(v T) {

}

func (q *Queue[T]) Size() int {
	return q.list.Size()
}

func (q *Queue[T]) IsEmpty() bool {
	return false
}
