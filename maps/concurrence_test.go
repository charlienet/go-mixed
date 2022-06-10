package maps

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
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

func BenchmarkLoadStore(b *testing.B) {
	ms := []Map[string, string]{
		NewRWMap[string, string](),
		NewConcurrentMap[string, string](),
		NewHashMap[string, string]().Synchronize(),
		NewHashMap[string, string](),
	}

	for _, m := range ms {
		var i int64
		b.Run(fmt.Sprintf("%T", m), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				gid := int(atomic.AddInt64(&i, 1) - 1)

				if gid == 0 {
					m.Set("0", strconv.Itoa(n))
				} else {
					m.Get("0")
				}
			}
		})
	}
}

func BenchmarkLoadStoreCollision(b *testing.B) {
	ms := []Map[string, string]{
		NewRWMap[string, string](),
		NewConcurrentMap[string, string](),
		NewHashMap[string, string]().Synchronize(),
		// &sync.Map{},
	}

	// 测试对于同一个 key 的 n-1 并发读和 1 并发写的性能
	for _, m := range ms {
		b.Run(fmt.Sprintf("%T", m), func(b *testing.B) {
			var i int64
			b.RunParallel(func(pb *testing.PB) {
				// 记录并发执行的 goroutine id
				gid := int(atomic.AddInt64(&i, 1) - 1)

				if gid == 0 {
					for i := 0; pb.Next(); i++ {
						m.Set("0", strconv.Itoa(i))
					}
				} else {
					for pb.Next() {
						m.Get("0")
					}
				}
			})
		})
	}
}
