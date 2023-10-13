package redis

import "testing"

func TestRename(t *testing.T) {
	New(&ReidsOption{
		Addrs: []string{"192.168.123.100:6379"},
	})
}
