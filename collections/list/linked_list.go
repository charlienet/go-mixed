package list

import (
	"github.com/charlienet/go-mixed/locker"
)

type LinkedList[T any] struct {
	list[T]
	front, tail *LinkedNode[T]
}

type LinkedNode[T any] struct {
	Value      T
	Prev, Next *LinkedNode[T]
}

// NewLinkedList 初始化链表
func NewLinkedList[T any](elems ...T) *LinkedList[T] {
	l :=
		&LinkedList[T]{
			list: list[T]{locker: locker.EmptyLocker},
		}

	for _, e := range elems {
		l.pushBackNode(&LinkedNode[T]{Value: e})
	}

	return l
}

func (l *LinkedList[T]) PushBack(v T) *LinkedList[T] {
	l.pushBackNode(&LinkedNode[T]{Value: v})

	return l
}

func (l *LinkedList[T]) PushFront(v T) *LinkedList[T] {
	l.pushFrontNode(&LinkedNode[T]{Value: v})

	return l
}

func (l *LinkedList[T]) FrontNode() *LinkedNode[T] {
	return l.front
}

func (l *LinkedList[T]) Front() T {
	return l.FrontNode().Value
}

func (l *LinkedList[T]) BackNode() *LinkedNode[T] {
	return l.tail
}

func (l *LinkedList[T]) Back() T {
	if l.size == 0 {
		panic(ErrorOutOffRange)
	}
	return l.tail.Value
}

func (l *LinkedList[T]) ForEach(fn func(T) bool) {
	l.locker.RLock()
	defer l.locker.RUnlock()

	for current := l.front; current != nil; current = current.Next {
		if fn(current.Value) {
			break
		}
	}
}

func (l *LinkedList[T]) GetAt(i int) T {
	if i <= l.Size() {
		var n int
		for current := l.front; current != nil; current = current.Next {
			if n == i {
				return current.Value
			}
			n++
		}
	}

	return *new(T)
}

func (l *LinkedList[T]) Remove(n *LinkedNode[T]) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		l.tail = n.Prev
	}

	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		l.front = n.Next
	}

	n.Next = nil
	n.Prev = nil

	l.size--
}

func (l *LinkedList[T]) RemoveAt(index int) {
	l.locker.Lock()
	defer l.locker.Unlock()

	var i int
	for current := l.front; current != nil; current = current.Next {
		if i == index {

			// 重连接
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev

			current.Prev = nil
			current.Next = nil

			l.size--
			break
		}

		i++
	}
}

func (l *LinkedList[T]) pushBackNode(n *LinkedNode[T]) {
	l.locker.Lock()
	defer l.locker.Unlock()

	n.Next = nil
	n.Prev = l.tail

	if l.tail != nil {
		l.tail.Next = n
	} else {
		l.front = n
	}

	l.tail = n

	l.size++
}

func (l *LinkedList[T]) pushFrontNode(n *LinkedNode[T]) {
	l.locker.Lock()
	defer l.locker.Unlock()

	n.Next = l.front
	n.Prev = nil
	if l.front != nil {
		l.front.Prev = n
	} else {
		l.tail = n
	}
	l.front = n

	l.size++
}
