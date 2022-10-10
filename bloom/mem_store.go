package bloom

import "github.com/bits-and-blooms/bitset"

type memStore struct {
	size uint
	set  *bitset.BitSet // 内存位图
}

func newMemStore(size uint) *memStore {
	return &memStore{
		size: size,
		set:  bitset.New(size),
	}
}

func (s *memStore) Clear() {
	s.set.ClearAll()
}

func (s *memStore) Set(offsets ...uint) error {
	for _, p := range offsets {
		s.set.Set(p)
	}

	return nil
}

func (s *memStore) Test(offsets ...uint) (bool, error) {
	for _, p := range offsets {
		if !s.set.Test(p) {
			return false, nil
		}
	}

	return true, nil
}
