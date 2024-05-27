package locker_test

import (
	"sync"
	"testing"
	"time"

	"github.com/charlienet/go-mixed/locker"
	"github.com/stretchr/testify/assert"
)

var sourcekey = "u-0001"

func TestTryLock(t *testing.T) {
	l := locker.NewSourceLocker()
	l.Lock("aa")

	assert.False(t, l.TryLock("aa"))
	assert.True(t, l.TryLock("bb"))

	defer l.Unlock("aa")
}

func TestM(t *testing.T) {
	l := locker.NewSourceLocker()

	for i := 0; i < 10000000; i++ {
		l.Lock("aaa")
		l.Unlock("aaa")
	}

	t.Logf("%+v", l)
}

func TestSourceLocker(t *testing.T) {
	l := locker.NewSourceLocker()

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

	l := locker.NewSourceLocker()

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
	l := locker.NewSourceLocker()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			l.Lock(sourcekey)
			l.Unlock(sourcekey)
		}
	})
}
