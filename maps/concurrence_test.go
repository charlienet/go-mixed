package maps

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConcurrence(t *testing.T) {
	m := NewHashMap(map[string]any{"a": "a"}).Synchronize()

	var wg = &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			m.Set("b", "b")
			m.Get("a")

			time.Sleep(time.Second)
		}()
	}

	wg.Wait()
}

func TestConcurrenceMapForeach(t *testing.T) {
	m := NewConcurrentMap[string, string]()
	m.Set("a", "a")

	m.ForEach(func(s1, s2 string) bool {
		fmt.Println(s1, s2)
		return false
	})
}

func BenchmarkItor(b *testing.B) {
	m := NewHashMap[string, string]().Synchronize()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Set("a", "b")
			for entity := range m.Iter() {
				m.Delete(entity.Key)
			}
		}
	})

}

func BenchmarkMap(b *testing.B) {
	m := NewHashMap[string, string]().Synchronize()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Set("a", "a")

			m.ForEach(func(s1, s2 string) bool {
				m.Delete("a")
				return false
			})
		}
	})
}
