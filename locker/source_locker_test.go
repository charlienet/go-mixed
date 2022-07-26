package locker

import (
	"sync"
	"testing"
	"time"
)

var sourcekey = "u-0001"

func TestTryLock(t *testing.T) {

}

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
	t.Logf("%+v", l)
}

func TestSourceTryLock(t *testing.T) {
	c := 5
	n := 0
	wg := new(sync.WaitGroup)
	wg.Add(c)

	l := NewSourceLocker()

	for i := 0; i < c; i++ {
		go func() {
			defer wg.Done()
			if l.TryLock(sourcekey) {
				n++
				time.Sleep(time.Second)

				l.Unlock(sourcekey)
			}
		}()
	}

	wg.Wait()
	t.Log("n:", n)
	t.Logf("%+v", l)
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
