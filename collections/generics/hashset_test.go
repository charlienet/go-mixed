package generics_test

import (
	"testing"

	"github.com/charlienet/go-mixed/collections/generics"
)

func TestSet(t *testing.T) {

	s := generics.NewHashSet[int]()
	s.Add(1, 2, 3)

	expected := generics.NewHashSet(1, 2, 3)

	_ = expected
}
