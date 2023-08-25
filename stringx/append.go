package stringx

import "fmt"

func Append[E any](dst []byte, e E) []byte {
	toAppend := fmt.Sprintf("%v", e)
	return append(dst, []byte(toAppend)...)
}
