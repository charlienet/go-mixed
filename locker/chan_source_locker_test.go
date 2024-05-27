package locker_test

import (
	"testing"

	"github.com/charlienet/go-mixed/locker"
)

func TestChanSourceLocker(t *testing.T) {
	l := locker.NewChanSourceLocker()
	c, ok := l.Get("aaaa")
	if ok {
		<-c

		println("ok")
	}

	println("fail")
}
