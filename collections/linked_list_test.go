package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkdList(t *testing.T) {
	l := NewLinkedList(1, 2, 3, 4, 5)

	l.Append(55)
	t.Log(l)

	t.Log(NewLinkedList("a", "b", "c", "d"))
}

func TestLinkedListSize(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.Append(5)
	ll.Append(6)
	ll.Append(7)

	assert.Equal(t, 3, ll.Size())

	ll.Prepend(8, 9, 10)
	assert.Equal(t, 6, ll.Size())

	ll.RemoveAt(1)
	assert.Equal(t, 5, ll.Size())
	t.Log(ll)

	ll.RemoveAt(ll.Size() - 1)
	t.Log(ll)
}
