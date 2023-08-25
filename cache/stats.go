package cache

import "sync/atomic"

type Stats struct {
	Hits   uint64
	Misses uint64
}

func (s *Stats) IncrementHits() {
	atomic.AddUint64(&s.Hits, 1)
}

func (s *Stats) IncrementMisses() {
	atomic.AddUint64(&s.Misses, 1)
}

func (c *Cache) Stats() *Stats {
	return c.stats
}
