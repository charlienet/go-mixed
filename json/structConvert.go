package json

import (
	"reflect"

	"github.com/charlienet/go-mixed/bytesconv"
)

// 结构转换为json字符串
func StructToJsonIndent(obj any) string {
	b, _ := MarshalIndent(obj, "", "  ")
	return bytesconv.BytesToString(b)
}

// 结构转换为json字符串
func StructToJson(obj any) string {
	b, _ := Marshal(obj)
	return bytesconv.BytesToString(b)
}

func StructToMap(obj any) map[string]any {
	typ := reflect.TypeOf(obj)

	kind := typ.Kind()
	if kind == reflect.Map {
		return toMap(obj)
	}

	val := reflect.ValueOf(obj)

	m := make(map[string]any)
	for i := 0; i < val.NumField(); i++ {
		m[typ.Field(i).Name] = val.Field(i).Interface()
	}

	return m
}

func StructToMapViaJson(obj any) map[string]any {
	m := make(map[string]any)

	j, _ := Marshal(obj)
	_ = Unmarshal(j, &m)

	return m
}

func toMap(obj any) map[string]any {
	if h, ok := obj.(map[string]any); ok {
		return h
	}

	return StructToMapViaJson(obj)
}
