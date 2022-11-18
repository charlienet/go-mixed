package cache

import "sync/atomic"

type Stats struct {
	Hits   uint64
	Misses uint64
}

func (s *Stats) AddHits() {
	atomic.AddUint64(&s.Hits, 1)
}

func (s *Stats) AddMisses() {
	atomic.AddUint64(&s.Misses, 1)
}

func (c *Cache) Stats() *Stats {
	return c.stats
}
