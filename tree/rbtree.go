package tree

type color uint

const (
	RED   color = 0
	BLACK color = 1
)

type RBTree[T Item[T]] struct {
	root *redBlackNode[T]
}

// 初始化红黒树
func NewRBTree[T Item[T]]() *RBTree[T] {
	return &RBTree[T]{}
}

func (t *RBTree[T]) ReplaceOrInsert(item Item[T]) Item[T] {
	return item
}

func (t *RBTree[T]) Delete(item T) (T, bool) {
	return *new(T), false
}

func (t *RBTree[T]) Get(key T) (T, bool) {
	return *new(T), false
}

func (t *RBTree[T]) Has(key T) bool {
	return false
}

// 红黒树节点
type redBlackNode[T Item[T]] struct {
	Color  color            // 颜色
	Item   T                // 数据
	parent *redBlackNode[T] // 父节点
	left   *redBlackNode[T] // 左
	right  *redBlackNode[T] // 右
}
