package collections

import "github.com/charlienet/go-mixed/locker"

type options struct {
	mu locker.RWLocker
}

func emptyLocker() locker.RWLocker {
	return locker.EmptyLocker
}
