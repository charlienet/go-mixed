package bytesconv

import (
	"fmt"
	"unsafe"
)

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type endian int

const (
	BigEndian endian = iota + 1
	LittleEndian
)

func (e endian) BytesToUInt64(data []byte) (uint64, error) {
	if len(data) > 8 {
		return 0, fmt.Errorf("bytes to uint64, bytes length is invaild")
	}

	var ret uint64
	var len int = len(data)

	if e == BigEndian {
		for i := 0; i < len; i++ {
			ret = ret | (uint64(data[len-1-i]) << (i * 8))
		}
	} else {
		for i := 0; i < len; i++ {
			ret = ret | (uint64(data[i]) << (i * 8))
		}
	}

	return ret, nil
}

func BytesToUInt64Big() {
	// binary.BigEndian
	// binary.LittleEndian
}
