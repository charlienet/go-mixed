package idgenerator

import (
	"sync"
	"time"

	"github.com/charlienet/go-mixed/idGenerator/store"
)

type obtainFunc func() (*store.Segment, error)

type doubleBuffer struct {
	current  *store.Segment // 当前
	backup   *store.Segment // 备用
	obtain   obtainFunc     // 数据段获取函数
	inFill   bool           // 备用缓冲填充中
	isReadly bool           // 备用缓冲区填充完成
	mu       sync.Mutex
}

func newDoubleBuffer(obtainFunc obtainFunc) *doubleBuffer {
	b := &doubleBuffer{obtain: obtainFunc}
	b.current, _ = b.obtain()

	return b
}

func (b *doubleBuffer) allot() (int64, bool) {
	if !b.inFill && b.current.IsEnding() {
		go b.full() // 填充备用缓冲
	}

	// 缓冲区耗尽时切换
	if b.current.IsEmpty() {
		// 检查备用缓冲是否已经填充完成，已完成时切换，否则等待
		for !b.isReadly {
			time.Sleep(time.Microsecond * 100)
		}

		b.switchBuf()
	}

	return b.current.Allot(), b.current.Reback()
}

func (b *doubleBuffer) full() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.inFill {
		var err error
		b.inFill = true

		b.backup, err = b.obtain()
		if err != nil {
			println("填充失败:", err.Error())
			panic(err)
		}
		b.isReadly = true
	}
}

func (b *doubleBuffer) switchBuf() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.isReadly {
		b.current, b.backup = b.backup, b.current
		b.inFill = false
		b.isReadly = false
	}
}
