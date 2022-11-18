package list_test

import (
	"testing"

	"github.com/charlienet/go-mixed/collections/list"
	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.ForEach(func(i int) bool {
		t.Log(i)
		return false
	})
}

func TestPushFront(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	l.ForEach(func(i int) bool {
		t.Log(i)
		return false
	})
}

func TestRemoveAt(t *testing.T) {

	l := list.NewLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.RemoveAt(1)

	l.ForEach(func(i int) bool {
		t.Log(i)
		return false
	})

	t.Log()

	l.RemoveAt(0)
	l.ForEach(func(i int) bool {
		t.Log(i)
		return false
	})

}

func TestSize(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	assert.Equal(t, 3, l.Size())

	l.RemoveAt(0)
	assert.Equal(t, 2, l.Size())
}

func TestLinkedListToSlice(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	s := l.ToSlice()
	t.Log(s)
}

func BenchmarkLinkedList(b *testing.B) {
	l := list.NewLinkedList[int]()
	l.Synchronize()

	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}
}

func TestRemoveNode(t *testing.T) {
	l := list.NewLinkedList(1, 2, 4)

	l.ForEach(func(i int) bool {
		t.Log(i)

		return false
	})

}
