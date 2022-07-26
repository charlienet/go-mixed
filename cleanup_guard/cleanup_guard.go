package cleanupguard

import "sync"

type CleanupGuard struct {
	enable bool
	fn     func()
	mutex  sync.Mutex
}

// 新建清理
func NewCleanupGuard(fn func()) CleanupGuard {
	return CleanupGuard{fn: fn, enable: true}
}

func (g *CleanupGuard) Enable() {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.enable = true
}

func (g *CleanupGuard) Run() {
	g.fn()
}
