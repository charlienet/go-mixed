package bytesconv

import (
	"bytes"
	"encoding/gob"
)

func Encode(v any) ([]byte, error) {
	var buf = new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(b []byte, out any) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	return dec.Decode(out)
}
