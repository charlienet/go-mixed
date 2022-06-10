package locker

import (
	"sync"
	"testing"
)

func TestSpinLock(t *testing.T) {
	l := NewSpinLocker()

	n := 10
	c := 0

	wg := new(sync.WaitGroup)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			l.Lock()
			c++
			l.Unlock()
		}()
	}

	wg.Wait()

	l.Lock()
	t.Log(c)
	l.Unlock()
}
