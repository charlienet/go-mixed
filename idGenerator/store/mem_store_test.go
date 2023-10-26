package store_test

import (
	"testing"

	"github.com/charlienet/go-mixed/idGenerator/store"
)

func TestMemSmall(t *testing.T) {
	s := store.NewMemStore(2)
	for i := 0; i < 10; i++ {
		t.Log(s.Assign(1, 9, 20))
	}
}

func TestMemBig(t *testing.T) {
	s := store.NewMemStore(2)
	for i := 0; i < 10; i++ {
		t.Log(s.Assign(0, 99, 18))
	}
}
