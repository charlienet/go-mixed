package tree

type avlTree struct {
}

func NewAVLTree() {

}

// 左单旋，新插入的节点在右子树的右侧
func (t *avlTree) rotateL() {

}

// 右单旋，新插入的节点在左子树的左侧
func (t *avlTree) rotateR() {

}

// 右左双旋，新插入的节点在右子树的左侧
// 1. 先对subR进行一个右单旋
// 2. 再对parent进行一个左单旋然后修改平衡因子
func (t *avlTree) rotateRL() {

}
