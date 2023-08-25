package ketama

import (
	"github.com/charlienet/go-mixed/locker"
	"github.com/charlienet/go-mixed/maps"
)

type Ketama struct {
	mu       locker.WithRWLocker
	replicas int
	m        maps.Map[uint64, string]
}

func New() *Ketama {
	return &Ketama{
		m: maps.NewHashMap[uint64, string](),
	}
}

func (k *Ketama) Synchronize() {
	k.mu.Synchronize()
}

func (k *Ketama) Add(nodes ...string) {
	k.mu.Lock()
	defer k.mu.Unlock()

	for _, node := range nodes {
		_ = node
	}
}

func (k *Ketama) IsEmpty() bool {
	k.mu.RLock()
	defer k.mu.RUnlock()

	return k.m.Count() == 0
}
