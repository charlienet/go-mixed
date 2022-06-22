package snowflake

import "testing"

func TestGetId(t *testing.T) {
	s := CreateSnowflake(22)
	for i := 0; i < 100; i++ {
		t.Log(s.GetId())
	}
}
