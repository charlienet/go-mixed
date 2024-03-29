//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

func RegisterFuzzyDecoders() {
}

var (
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)
