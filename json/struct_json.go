package json

import "github.com/charlienet/go-mixed/bytesconv"

func StructToJsonIndent(obj interface{}) string {
	b, _ := MarshalIndent(obj, "", "  ")
	return bytesconv.BytesToString(b)
}

// 结构转换为json字符串
func StructToJson(obj interface{}) string {
	b, _ := Marshal(obj)
	return bytesconv.BytesToString(b)
}
