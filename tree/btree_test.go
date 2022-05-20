package tree

import "testing"

func TestNewBTree(t *testing.T) {
	tr := NewBTree[Int](32)

	_ = tr
}

type Int int

func (a Int) Less(b Int) bool {
	return a < b
}
