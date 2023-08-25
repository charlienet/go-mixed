package rbtree

type color bool

const (
	black, red color = true, false
)

type TreeNode[K any, V any] struct {
}
