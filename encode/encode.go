package encode

import "encoding/hex"

type BytesResult []byte

func FromString(s string) BytesResult {
	return BytesResult([]byte(s))
}

func FromHexString(s string) BytesResult {
	b, _ := hex.DecodeString(s)
	return BytesResult(b)
}

func FromBytes(b []byte) BytesResult {
	return BytesResult(b)
}

