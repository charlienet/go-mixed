package sets_test

import (
	"fmt"
	"testing"

	"github.com/charlienet/go-mixed/json"
	"github.com/charlienet/go-mixed/sets"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	s := sets.NewHashSet("abc", "bcd")
	t.Log(s)
}

func SortedSetStrign(t *testing.T) {
	s := sets.NewSortedSet("abc", "bcd")
	t.Log(s)
}

func TestContains(t *testing.T) {
	s := sets.NewHashSet("abc", "bcd", "efg", "b")
	assert.Equal(t, true, s.Contains("b"))
}

func TestContainsAll(t *testing.T) {

}

func TestContainsAny(t *testing.T) {

}

func TestMarshal(t *testing.T) {
	s := sets.NewHashSet("abc", "bcd", "efg", "b")
	t.Log(json.StructToJsonIndent(s))
}

type GenericSet[T any] interface {
	WithField(GenericSet[T]) T
	Info(args ...any)
}

var _ GenericSet[*MySet[string]] = &MySet[string]{}

type MySet[T comparable] struct {
}

func (m *MySet[T]) WithField(other GenericSet[*MySet[T]]) *MySet[T] {
	other.Info("with field")

	return m
}

func (m *MySet[T]) Info(args ...any) {
	fmt.Println("abc", args)
}

func DoStuff[T GenericSet[T]](t T) {
	t.WithField(t).Info("here")
}

func TestDoStuff(t *testing.T) {
	DoStuff(&MySet[string]{})

	// DoStuff(sets.NewHashSet("aaa"))
}

func TestUnion(t *testing.T) {
	ret := sets.Union[string](sets.NewHashSet("abc", "bcd", "e"), sets.NewHashSet("abc", "f", "bcd"))
	t.Log(ret)
}

func TestDifference(t *testing.T) {
	ret := sets.Difference[string](sets.NewHashSet("abc", "bcd", "e"), sets.NewHashSet("abc", "f", "bcd"))
	t.Log(ret)
}

func TestIntersection(t *testing.T) {
	ret := sets.Intersection[string](sets.NewHashSet("abc", "bcd", "e"), sets.NewHashSet("abc", "f", "bcd"))
	t.Log(ret)
}
