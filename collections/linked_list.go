package collections

import (
	"fmt"

	"github.com/charlienet/go-mixed/locker"
)

// 单链表节点
type linkedNode[T any] struct {
	item T
	next *linkedNode[T]
}

// 双向链接节点
type doubleLinkedNode[T any] struct {
	item T
	prev *linkedNode[T]
	next *linkedNode[T]
}

type LinkedList[T any] struct {
	mu         locker.RWLocker
	head, tail *linkedNode[T]
	size       int
}

// 初始化单向链表
func NewLinkedList[T any](vals ...T) *LinkedList[T] {
	l := &LinkedList[T]{mu: locker.EmptyLocker}
	l.Append(vals...)

	return l
}

// synchronized
func (l *LinkedList[T]) Synchronize() {
	l.mu = locker.NewRWLocker()
}

func (l *LinkedList[T]) Append(vals ...T) {
	for _, v := range vals {
		l.append(v)
	}
}

// prepend will prepend the list with a value, the reference node is Returned
func (l *LinkedList[T]) Prepend(vals ...T) {
	for _, v := range vals {
		l.prepend(v)
	}
}

func (l *LinkedList[T]) append(v T) {
	n := createSingleNode(v)

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head == nil {
		l.head = n
	}

	if l.tail != nil {
		l.tail.next = n
	}

	l.tail = n
	l.size++
}

func (l *LinkedList[T]) prepend(v T) {
	n := createSingleNode(v)

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.head != nil {
		n.next = l.head
	}

	if l.tail == nil {
		l.tail = n
	}

	l.head = n
	l.size++
}

func (l *LinkedList[T]) PushHead() {
}

func (l *LinkedList[T]) PushTail() {

}

func (l *LinkedList[T]) Exist(v T) bool {
	ret := false
	l.ForEach(func(t T) bool {
		// if t == v {
		// 	ret = true
		// 	return true
		// }

		return false
	})

	return ret
}

func (l *LinkedList[T]) RemoveAt(index int) {
	var i int
	var prev *linkedNode[T]

	prev = l.head
	for current := l.head; current != nil; {

		if i == index {
			prev.next = current.next
			current.next = nil

			l.size--
			return
		}

		prev = current
		current = current.next
		i++
	}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) GetAt(i int) T {
	if i <= l.Size() {
		var n int
		for current := l.head; current != nil; current = current.next {
			if n == i {
				return current.item
			}
			n++
		}
	}

	return *new(T)
}

func (l *LinkedList[T]) ForEach(fn func(T) bool) {
	for current := l.head; current != nil; current = current.next {
		if fn(current.item) {
			break
		}
	}
}

func (l *LinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[T]) ToList() []T {
	ret := make([]T, 0, l.size)
	l.ForEach(func(t T) bool {
		ret = append(ret, t)
		return false
	})

	return ret
}

func (l *LinkedList[T]) String() string {
	return fmt.Sprint(l.ToList())
}

func (l *LinkedList[T]) remove(n *linkedNode[T]) {
	n.next = nil
	l.size--
}

type DLinkedList[T any] struct {
	head, tail *doubleLinkedNode[T]
	size       int
}

// 初始化双向链表
func NewDoubleLinkedLis[T any]() *DLinkedList[T] {
	return &DLinkedList[T]{}
}

func createSingleNode[T any](v T) *linkedNode[T] {
	return &linkedNode[T]{item: v}
}

func createDoubleNode[T any](v T) *doubleLinkedNode[T] {
	return &doubleLinkedNode[T]{item: v}
}
