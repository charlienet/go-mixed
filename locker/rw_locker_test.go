package locker

import "testing"

func TestRWLokcer(t *testing.T) {
	l := NewRWLocker()
	l.RLock()

	t.Log(l.TryRLock())

	l.RUnlock()
}
