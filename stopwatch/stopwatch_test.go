package stopwatch_test

import (
	"testing"
	"time"
	_ "unsafe"

	"github.com/charlienet/go-mixed/stopwatch"
)

func TestWatch(t *testing.T) {
	watch := stopwatch.StartNew()

	time.Sleep(time.Second * 3)
	t.Log("Elapsed:", watch.Elapsed())
	t.Log("Elapsed:", watch.ElapsedMilliseconds())
	t.Log("Elapsed:", watch.ElapsedMicroseconds())
	t.Log("Elapsed:", watch.ElapsedNanoseconds())

	time.Sleep(time.Second * 1)
	t.Log("Elapsed:", watch.Elapsed())

	watch.Restart()
	t.Log("Elapsed:", watch.Elapsed())
	time.Sleep(time.Second * 1)
	t.Log("Elapsed:", watch.Elapsed())

	watch.Reset()

	watch.Restart()
}

func TestStop(t *testing.T) {
	watch := stopwatch.StartNew()
	time.Sleep(time.Second)

	watch.Stop()
	time.Sleep(time.Second)
	watch.Start()

	time.Sleep(time.Second)

	t.Log(watch.Elapsed())
}
