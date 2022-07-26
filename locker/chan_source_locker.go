package locker

type ChanLocker interface {
	Get(key string) (ch <-chan int, ok bool)
	Release(key string)
}

type chanSourceLock struct {
	m       RWLocker
	content map[string]chan int
}

func (s *chanSourceLock) Get(key string) (ch <-chan int, ok bool) {
	s.m.RLock()
	ch, ok = s.content[key]
	s.m.RUnlock()
	if ok {
		return
	}
	s.m.Lock()
	ch, ok = s.content[key]
	if ok {
		s.m.Unlock()
		return
	}
	s.content[key] = make(chan int)
	ch = s.content[key]
	ok = true
	s.m.Unlock()
	return
}

func (s *chanSourceLock) Release(key string) {
	s.m.Lock()
	ch, ok := s.content[key]
	if ok {
		close(ch)
		delete(s.content, key)
	}
	s.m.Unlock()
}
