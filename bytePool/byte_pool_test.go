package bytecache

import "testing"

func TestByteCache(t *testing.T) {
	bp := NewBytePool(512, 1024, 1024)
	buffer := bp.Get()
	defer bp.Put(buffer)

	t.Log(len(buffer))
}
