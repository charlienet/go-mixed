package tree

// B+ tree

type BPTree[T Item[T]] struct {
	length int
	degree int
}

// 初始化B+树
func NewBPTree[T Item[T]](degree int) *BPTree[T] {
	return &BPTree[T]{
		degree: degree,
	}
}

// nil cannot be added to the tree (will panic).
func (t *BPTree[T]) ReplaceOrInsert(item Item[T]) Item[T] {
	return item
}

func (t *BPTree[T]) Delete(item T) (T, bool) {
	return *new(T), false
}

func (t *BPTree[T]) Get(key T) (T, bool) {
	return *new(T), false
}

// Has returns true if the given key is in the tree.
func (t *BPTree[T]) Has(key T) bool {
	return false
}

// 升序
func (t *BPTree[T]) Ascend(iterator ItemIterator[T]) {
}

// 大于等于
func (t *BPTree[T]) AscendGreaterOrEqual(pivot Item[T], iterator ItemIterator[T]) {

}

// 小于
func (t *BPTree[T]) AscendLessThan(pivot Item[T], iterator ItemIterator[T]) {

}

// 范围迭代
func (t *BPTree[T]) AscendRange(greaterOrEqual, lessThan Item[T], iterator ItemIterator[T]) {

}

// 降序升序迭代
func (t *BPTree[T]) Descend(iterator ItemIterator[T]) {
}

// 大于，降序迭代
func (t *BPTree[T]) DescendGreaterThan(pivot Item[T], iterator ItemIterator[T]) {

}

// 小于等于，降序迭代
func (t *BPTree[T]) DescendLessOrEquql(pivot Item[T], iterator ItemIterator[T]) {

}

// 范围降序迭代
func (t *BPTree[T]) DescendRange(greaterOrEqual, lessThan Item[T], iterator ItemIterator[T]) {

}

func (t *BPTree[T]) Min() (T, bool) {
	return *new(T), false
}

func (t *BPTree[T]) Max() (T, bool) {
	return *new(T), false
}

func (t *BPTree[T]) Len() int {
	return t.length
}

func (t *BPTree[T]) Clear() {
}

// B+ 树节点
type bpNode[T Item[T]] struct {
}
