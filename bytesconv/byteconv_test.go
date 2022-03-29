package bytesconv

import (
	"testing"
)

func TestBytesToUint64(t *testing.T) {
	t.Log(BigEndian.BytesToUInt64([]byte{0x88, 0x45}))
	t.Log(LittleEndian.BytesToUInt64([]byte{0x88, 0x45}))
}

func BenchmarkBytesToUInt64(b *testing.B) {
	aa := []byte("a")

	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BigEndian.BytesToUInt64(aa)
		}
	})
}
