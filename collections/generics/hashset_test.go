package generics_test

import (
	"testing"

	"github.com/charlienet/go-mixed/collections/generics"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {

	s := generics.NewHashSet[int]()
	s.Add(1, 2, 3)

	expected := generics.NewHashSet(1, 2, 3)

	_ = expected
}

func TestSetExist(t *testing.T) {
	s := generics.NewHashSet(1, 2, 3)
	assert.True(t, s.Contain(1))
	assert.False(t, s.Contain(5))
}
