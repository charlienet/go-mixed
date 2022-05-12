package maps

import (
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
