package tree

import "sync"

const (
	DefaultFreeListSize = 32
)

type FreeList[T Item[T]] struct {
	mu       sync.Mutex
	freelist []*bNode[T]
}

// 初始化B树
type BTree[T Item[T]] struct {
	degree int
	length int
	root   *bNode[T]
}

func NewBTree[T Item[T]](degree int) *BTree[T] {
	return &BTree[T]{}
}

// nil cannot be added to the tree (will panic).
func (t *BTree[T]) ReplaceOrInsert(item Item[T]) Item[T] {
	return item
}

func (t *BTree[T]) Delete(item T) (T, bool) {
	return *new(T), false
}

func (t *BTree[T]) Get(key T) (T, bool) {
	return *new(T), false
}

// Has returns true if the given key is in the tree.
func (t *BTree[T]) Has(key T) bool {
	return false
}

// 升序
func (t *BTree[T]) Ascend(iterator ItemIterator[T]) {
}

// 大于等于
func (t *BTree[T]) AscendGreaterOrEqual(pivot Item[T], iterator ItemIterator[T]) {

}

// 小于
func (t *BTree[T]) AscendLessThan(pivot Item[T], iterator ItemIterator[T]) {

}

// 范围迭代
func (t *BTree[T]) AscendRange(greaterOrEqual, lessThan Item[T], iterator ItemIterator[T]) {

}

// 降序升序迭代
func (t *BTree[T]) Descend(iterator ItemIterator[T]) {
}

// 大于，降序迭代
func (t *BTree[T]) DescendGreaterThan(pivot Item[T], iterator ItemIterator[T]) {

}

// 小于等于，降序迭代
func (t *BTree[T]) DescendLessOrEquql(pivot Item[T], iterator ItemIterator[T]) {

}

// 范围降序迭代
func (t *BTree[T]) DescendRange(greaterOrEqual, lessThan Item[T], iterator ItemIterator[T]) {

}

func (t *BTree[T]) Min() (T, bool) {
	return *new(T), false
}

func (t *BTree[T]) Max() (T, bool) {
	return *new(T), false
}

func (t *BTree[T]) Len() int {
	return t.length
}

func (t *BTree[T]) Clear() {

}

// items stores items in a node.
type items[T Item[T]] []T

// children stores child nodes in a node.
type children[T Item[T]] []*bNode[T]

type copyOnWriteContext[T Item[T]] struct {
	freelist *FreeList[T]
}

// B树的内部节点
type bNode[T Item[T]] struct {
	items    items[T]
	children children[T]
	cow      *copyOnWriteContext[T]
}
