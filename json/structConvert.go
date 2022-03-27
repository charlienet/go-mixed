package json

import (
	"reflect"

	"github.com/charlienet/go-mixed/bytesconv"
)

// 结构转换为json字符串
func StructToJsonIndent(obj interface{}) string {
	b, _ := MarshalIndent(obj, "", "  ")
	return bytesconv.BytesToString(b)
}

// 结构转换为json字符串
func StructToJson(obj interface{}) string {
	b, _ := Marshal(obj)
	return bytesconv.BytesToString(b)
}

func StructToMap(obj interface{}) map[string]interface{} {
	typ := reflect.TypeOf(obj)

	kind := typ.Kind()
	if kind == reflect.Map {
		return toMap(obj)
	}

	val := reflect.ValueOf(obj)

	m := make(map[string]interface{})
	for i := 0; i < val.NumField(); i++ {
		m[typ.Field(i).Name] = val.Field(i).Interface()
	}

	return m
}

func StructToMapViaJson(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	j, _ := Marshal(obj)
	_ = Unmarshal(j, &m)

	return m
}

func toMap(obj interface{}) map[string]interface{} {
	if h, ok := obj.(map[string]interface{}); ok {
		return h
	}

	return StructToMapViaJson(obj)
}
