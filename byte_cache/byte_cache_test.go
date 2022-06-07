package bytecache

import "testing"

func TestByteCache(t *testing.T) {
	bp := NewBytePool(512, 1024, 1024)
	buffer := bp.Get()
	defer bp.Put(buffer)

	t.Log(len(buffer))
}

func BenchmarkByteCache(b *testing.B) {
	bp := NewBytePool(512, 1024, 1024)
	b.Run("aaaaa", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bp.Put(bp.Get())
		}
	})

	b.Run("bbbb", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]byte, 1024)

			a[0] = 1
		}
	})
}
