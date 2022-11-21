package list_test

import (
	"testing"

	"github.com/charlienet/go-mixed/collections/list"
)

func TestNewArrayList(t *testing.T) {
	l := list.NewArrayList(1, 2, 3)

	l.ForEach(func(i int) {
		t.Log(i)
	})

}

func TestArrayPushBack(t *testing.T) {
	l := list.NewArrayList[int]()

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.ForEach(func(i int) {
		t.Log(i)
	})
}

func TestArrayPushFront(t *testing.T) {
	l := list.NewArrayList[int]()

	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	l.PushBack(99)
	l.PushBack(88)

	l.ForEach(func(i int) {
		t.Log(i)
	})
}
