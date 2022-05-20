package collections

type Item[T any] interface {
	Less(T) bool
}

// 单链表节点
type linkedNode[T Item[T]] struct {
	item T
	next *linkedNode[T]
}

// 双向链接节点
type doubleLinkedNode[T Item[T]] struct {
	item T
	prev *linkedNode[T]
	next *linkedNode[T]
}

type LinkedList[T Item[T]] struct {
	head, tail *linkedNode[T]
	size       int
}

// 初始化单向链表
func NewLinkedList[T Item[T]]() *LinkedList[T] {
	return &LinkedList[T]{}
}

type DLinkedList[T Item[T]] struct {
	head, tail *doubleLinkedNode[T]
	size       int
}

// 初始化双向链表
func NewDoubleLinkedLis[T Item[T]]() *DLinkedList[T] {
	return &DLinkedList[T]{}
}
