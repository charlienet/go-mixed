package locker_test

import (
	"testing"

	"github.com/charlienet/go-mixed/locker"
)

func TestLocker(t *testing.T) {

	var l locker.Locker

	l.Synchronize()

	l.Lock()
	defer l.Unlock()
}

func TestNew(t *testing.T) {
	var a locker.RWLocker
	a.Synchronize()

}

func TestSpinLocker(t *testing.T) {
	var l locker.SpinLocker
	l.Synchronize()

	l.Lock()
	defer l.Unlock()
}

func TestRWLocker(t *testing.T) {
	var l locker.RWLocker
	l.Lock()
}

func TestPointLocker(t *testing.T) {
	l := locker.NewLocker()
	l.Lock()
	l.Lock()

	defer l.Unlock()
}
