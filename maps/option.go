package maps

import (
	"github.com/charlienet/go-mixed/locker"
)

type options struct {
	hasLocker bool
	mu        locker.RWLocker
}

func acquireDefaultOptions() *options {
	return &options{mu: locker.NewEmptyLocker()}
}
