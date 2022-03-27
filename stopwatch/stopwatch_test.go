package stopwatch_test

import (
	"testing"
	"time"

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
