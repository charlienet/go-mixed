package stack_test

import (
	"testing"

	"github.com/charlienet/go-mixed/collections/stack"
)

func TestStack(t *testing.T) {
	arrayStack := stack.NewArrayStack[string]()
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")

	t.Log("size:", arrayStack.Size())
	t.Log("pop:", arrayStack.Pop())
	t.Log("pop:", arrayStack.Pop())
	t.Log("size:", arrayStack.Size())
	arrayStack.Push("drag")
	t.Log("pop:", arrayStack.Pop())
	arrayStack.Push("test")
	s := arrayStack.Pop()
	t.Log(s)
}
