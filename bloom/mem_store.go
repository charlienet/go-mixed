package bloom

import (
	"context"
	"sync"

	"github.com/bits-and-blooms/bitset"
)

type memStore struct {
	size uint
	set  *bitset.BitSet // 内存位图
	lock sync.RWMutex   // 同步锁
}

func newMemStore(size uint) *memStore {
	return &memStore{
		size: size,
		set:  bitset.New(size),
	}
}

func (s *memStore) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set.ClearAll()
}

func (s *memStore) Set(ctx context.Context, offsets ...uint) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, p := range offsets {
		s.set.Set(p)
	}

	return nil
}

func (s *memStore) Test(offsets ...uint) (bool, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, p := range offsets {
		if !s.set.Test(p) {
			return false, nil
		}
	}

	return true, nil
}
