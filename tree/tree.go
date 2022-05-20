package tree

type ITree[T Item[T]] interface {
	ReplaceOrInsert(item Item[T]) Item[T]                                    // nil cannot be added to the tree (will panic).
	Delete(item T) (T, bool)                                                 // 删除节点
	Get(key T) (T, bool)                                                     // 获取节点
	Has(key T) bool                                                          // 节点在树中存在时返回true
	Ascend(iterator ItemIterator[T])                                         // 升序
	AscendGreaterOrEqual(pivot Item[T], iterator ItemIterator[T])            // 大于等于
	AscendLessThan(pivot Item[T], iterator ItemIterator[T])                  // 小于
	AscendRange(greaterOrEqual, lessThan Item[T], iterator ItemIterator[T])  // 范围迭代
	Descend(iterator ItemIterator[T])                                        // 降序升序迭代
	DescendGreaterThan(pivot Item[T], iterator ItemIterator[T])              // 大于，降序迭代
	DescendLessOrEquql(pivot Item[T], iterator ItemIterator[T])              // 小于等于，降序迭代
	DescendRange(greaterOrEqual, lessThan Item[T], iterator ItemIterator[T]) // 范围降序迭代
	Min() (T, bool)
	Max() (T, bool)
	Len() int
	Clear()
}

// 表示树中的单个对象。
type Item[T any] interface {
	Less(than T) bool
}

// 迭代函数，返回true时停止迭代
type ItemIterator[T Item[T]] func(i T) bool
