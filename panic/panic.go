package panic

import (
	"context"
	"fmt"
	"runtime/debug"
)

type Panic struct {
	R     any
	Stack []byte
}

func (p Panic) String() string {
	return fmt.Sprintf("%v\n%s", p.R, p.Stack)
}

type PanicGroup struct {
	panics chan Panic // 致命错误通知
	dones  chan int   // 协程完成通知
	jobs   chan int   // 并发数量
	jobN   int32      // 工作协程数量
}

func NewPanicGroup(maxConcurrent int) *PanicGroup {
	return &PanicGroup{
		panics: make(chan Panic, 8),
		dones:  make(chan int, 8),
		jobs:   make(chan int, maxConcurrent),
	}
}

func (g *PanicGroup) Go(f func()) *PanicGroup {
	g.jobN++

	go func() {
		g.jobs <- 1
		defer func() {
			<-g.jobs
			// go 语言只能在自己的协程中捕获自己的 panic
			// 如果不处理，整个*进程*都会退出
			if r := recover(); r != nil {
				g.panics <- Panic{R: r, Stack: debug.Stack()}
				// 如果发生 panic 就不再通知 Wait() 已完成
				// 不然就可能出现 g.jobN 为 0 但 g.panics 非空
				// 的情况，此时 Wait() 方法需要在正常结束的分支
				// 中再额外检查是否发生了 panic，非常麻烦
				return
			}

			g.dones <- 1
		}()
		
		f()
	}()

	return g
}

func (g *PanicGroup) Wait(ctx context.Context) error {
	if g.jobN == 0 {
		panic("no job to wait")
	}

	for {
		select {
		case <-g.dones: // 协程正常结束
			g.jobN--
			if g.jobN == 0 {
				return nil
			}
		case p := <-g.panics: // 协程有 panic
			panic(p)
		case <-ctx.Done():
			// 整个 ctx 结束，超时或者调用方主动取消
			// 子协程应该共用该 ctx，都会收到相同的结束信号
			// 不需要在这里再去通知各协程结束（实现起来也麻烦）
			return ctx.Err()
		}
	}
}
