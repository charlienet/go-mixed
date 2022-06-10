package locker

import (
	"sync"
	"testing"
)

var sourcekey = "u-0001"

func TestSourceLocker(t *testing.T) {
	l := NewSourceLocker()

	c := 5
	n := 0
	wg := new(sync.WaitGroup)
	wg.Add(c)

	for i := 0; i < c; i++ {
		go func() {
			defer wg.Done()

			l.Lock(sourcekey)
			n++
			l.Unlock(sourcekey)
		}()
	}

	wg.Wait()
	t.Log("n:", n)
}

func BenchmarkSourceLocker(b *testing.B) {
	l := NewSourceLocker()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			l.Lock(sourcekey)
			l.Unlock(sourcekey)
		}
	})
}
