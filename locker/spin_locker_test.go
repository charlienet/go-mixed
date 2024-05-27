package locker_test

import (
	"sync"
	"testing"

	"github.com/charlienet/go-mixed/locker"
)

func TestSpinLock(t *testing.T) {
	l := locker.NewSpinLocker()

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
