package bytesconv

import "testing"

func TestBytesToUint64(t *testing.T) {
	t.Log(BigEndian.BytesToUInt64([]byte{0x88, 0x45}))
	t.Log(LittleEndian.BytesToUInt64([]byte{0x88, 0x45}))
}
