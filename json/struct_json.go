package json

import "github.com/charlienet/go-mixed/bytesconv"

func StructToJsonIndent(obj any) string {
	b, _ := MarshalIndent(obj, "", "  ")
	return bytesconv.BytesToString(b)
}

// 结构转换为json字符串
func StructToJson(obj any) string {
	b, _ := Marshal(obj)
	return bytesconv.BytesToString(b)
}
