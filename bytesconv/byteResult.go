package bytesconv

import (
	"encoding/base64"
	"encoding/hex"
)

type BytesResult []byte

func (r BytesResult) Hex() string {
	return hex.EncodeToString(r)
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
