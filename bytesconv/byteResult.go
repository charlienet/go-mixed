package bytesconv

import (
	"encoding/base64"
	"encoding/hex"
)

const hextable = "0123456789ABCDEF"

type BytesResult []byte

func (r BytesResult) Hex() string {
	return hex.EncodeToString(r)
}

func (r BytesResult) UppercaseHex() string {
	dst := make([]byte, hex.EncodedLen(len(r)))
	j := 0
	for _, v := range r {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}

	return BytesToString(dst)
}

func (r BytesResult) Base64() string {
	return base64.StdEncoding.EncodeToString(r)
}

func (r BytesResult) Bytes() []byte {
	return r
}

func (r BytesResult) String() string {
	return r.Hex()
}
