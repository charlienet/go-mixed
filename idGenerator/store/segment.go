package store

import "fmt"

// 号段
type Segment struct {
	start   int64
	end     int64
	current int64
	reback  bool
}

func (s *Segment) Allot() int64 {
	s.current++
	return s.current
}

func (s *Segment) IsEnding() bool {
	return (s.current - s.start) > (s.end - s.current)
}

func (s *Segment) IsEmpty() bool {
	return s.current == s.end
}

func (s *Segment) Reback() bool {
	// println("回旋确认:", s.reback, s.current == (s.start+1))
	return s.reback && s.current == (s.start+1)
}

func (s *Segment) String() string {
	return fmt.Sprintf("start:%d-%d(%v)", s.start, s.end, s.reback)
}
