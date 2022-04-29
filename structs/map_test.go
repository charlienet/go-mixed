package structs_test

import (
	"testing"

	"github.com/charlienet/go-mixed/structs"
)

func TestStructToMap(t *testing.T) {
	o := struct {
		UserName  string
		InTagName string `json:"in_tag_name,omitempty"`
		Ignore    string `json:"-"`
		KeepEmpty int
		OmitEmpty int `json:",omitempty"`
	}{
		UserName:  "测试字段",
		InTagName: "具体名称",
		Ignore:    "这个字段跳过",
		KeepEmpty: 0,
		OmitEmpty: 0,
	}

	t.Log(structs.Struct2Map(o, structs.TagName("struct")))
	t.Log(structs.Struct2Map(o, structs.IgnoreEmpty()))
	t.Log(structs.Struct2Map(o, structs.Omitempty()))
	t.Log(structs.Struct2Map(o, structs.Lcfirst()))
	t.Log(structs.Struct2Map(o, structs.Camel2Case()))
}

func TestMapToStruct(t *testing.T) {

}

func TestMap2Map(t *testing.T) {
	source := map[string]any{
		"Abc": 143,
	}

	structs.New(source)
}
