package locker

import "sync"

var locks = make(map[string]sync.Locker)

func Lock(name string) {
	if l, ok := locks[name]; ok {
		l.Lock()
	}

	new := &sync.Mutex{}
	locks[name] = new
	new.Lock()
}

func Unlock(name string) {
	if l, ok := locks[name]; ok {
		l.Unlock()
	}
}
