package json_test

import (
	"testing"

	"github.com/charlienet/go-mixed/json"
)

func TestNameConvert(t *testing.T) {
	d := struct {
		UserName string
		Age      int
	}{UserName: "测试", Age: 13}

	t.Log(json.StructToJsonIndent(json.CamelCase{Value: d}))
	t.Log(json.StructToJsonIndent(json.SnakeCase{Value: d}))
}
