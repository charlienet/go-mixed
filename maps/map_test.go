package maps_test

import (
	"testing"

	"github.com/charlienet/go-mixed/maps"
)

func TestMap(t *testing.T) {
	_ = maps.NewHashMap[string, any]()
}

func TestHashMap(t *testing.T) {
	var m maps.Map[string, any] = maps.NewHashMap[string, any]()

	_ = m

}

func TestIter(t *testing.T) {
	m := maps.NewHashMap[string, string]()
	m.Set("abc", "abc")

	for e := range m.Iter() {
		t.Log(e.Key, e.Value)
	}
}
