package stopwatch

import (
	"sync/atomic"
	"time"

	_ "unsafe"
)

type watch struct {
	startTime int64
	elapsed   int64
	isRunning bool
	count     int32
}

func StartNew() *watch {
	w := &watch{}
	w.Start()

	return w
}

func (w *watch) Start() {
	if !w.isRunning {
		w.startTime = w.RuntimeNano()
		w.isRunning = true
	}
}

// 将运行时间重置为零，并开始测量运行时间。
func (w *watch) Restart() {
	w.elapsed = 0
	w.startTime = w.RuntimeNano()
	w.isRunning = true
}

// 停止时间间隔测量并将经过的时间重置为零。
func (w *watch) Reset() {
	w.elapsed = 0
	w.startTime = w.RuntimeNano()
	w.isRunning = false
}

func (w *watch) Stop() {
	if w.isRunning {
		w.elapsed += w.RuntimeNano() - w.startTime
		w.isRunning = false

		if w.elapsed < 0 {
			w.elapsed = 0
		}
	}
}

func (w *watch) IsRunning() bool {
	return w.isRunning
}

func (w *watch) Elapsed() time.Duration {
	return time.Duration(w.ElapsedNanoseconds())
}

func (w *watch) ElapsedMilliseconds() int64 {
	return w.ElapsedNanoseconds() / 1e6
}

func (w *watch) ElapsedMicroseconds() int64 {
	return w.ElapsedNanoseconds() / 1e3
}

func (w *watch) ElapsedNanoseconds() int64 {
	num := w.elapsed
	if w.isRunning {
		num += w.RuntimeNano() - w.startTime
	}

	return num
}

func (w *watch) RuntimeNano() int64 {
	atomic.AddInt32(&w.count, 1)
	return runtimeNano()
}

//go:linkname runtimeNano runtime.nanotime
func runtimeNano() int64
