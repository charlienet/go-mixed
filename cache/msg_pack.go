package cache

import (
	"github.com/vmihailenco/msgpack/v5"
)

func Marshal(v any) ([]byte, error) {
	switch v := v.(type) {
	case nil:
		return nil, nil
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	}

	b, err := msgpack.Marshal(v)
	if err != nil {
		return nil, err
	}

	return b, err
}

func Unmarshal(b []byte, v any) error {
	if len(b) == 0 {
		return nil
	}

	switch v := v.(type) {
	case nil:
		return nil
	case *[]byte:
		clone := make([]byte, len(b))
		copy(clone, b)
		*v = clone
	case *string:
		*v = string(b)
		return nil
	}

	return msgpack.Unmarshal(b, v)
}
