package structs_test

import (
	"testing"

	"github.com/charlienet/go-mixed/structs"
)

func TestStructToMap(t *testing.T) {
	o := struct {
		Abc       string
		InTagName string `json:"in_tag_name,omitempty"`
		KeepEmpty int
		OmitEmpty int `json:",omitempty"`
	}{
		Abc:       "测试字段",
		InTagName: "具体名称",
		KeepEmpty: 0,
		OmitEmpty: 0,
	}

	t.Log(structs.ToMap(o))
	t.Log(structs.ToMap(o, structs.IgnoreEmpty()))
	t.Log(structs.ToMap(o, structs.Omitempty()))
}
