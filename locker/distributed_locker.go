package locker

import "context"

type DistributedLocker interface {
	Unlock(context.Context, string)
}
