package maps

import (
	"sync"

	"github.com/charlienet/go-mixed/locker"
)

type options struct {
	mu sync.Locker
}

func acquireDefaultOptions() *options {
	return &options{mu: locker.NewEmptyLocker()}
}
