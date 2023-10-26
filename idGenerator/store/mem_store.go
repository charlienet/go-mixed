package store

import (
	"sync"

	"github.com/charlienet/go-mixed/mathx"
)

type memStore struct {
	mu      sync.Mutex
	machine int64
	current int64
}

func NewMemStore(machineCode int64) *memStore {
	return &memStore{machine: machineCode}
}

func (s *memStore) UpdateMachineCode(max int64) (int64, error) {
	return s.machine, nil
}
func (s *memStore) MachineCode() int64 {
	return s.machine
}

func (s *memStore) Assign(min, max, step int64) (*Segment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	step = mathx.Min(step, max)
	start := mathx.Max(s.current, min)
	end := start + step

	reback := false
	if start >= max {
		start = min
		end = step
		s.current = end

		reback = true
	}

	if end > max {
		end = max
		s.current = end
	}

	s.current = end

	return &Segment{
		start:   start,
		current: start,
		end:     end,
		reback:  reback,
	}, nil
}

func (s *memStore) Close() {

}
