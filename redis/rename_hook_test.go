package redis

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRename(t *testing.T) {
	New(&RedisOption{
		Addrs: []string{"192.168.123.100:6379"},
	})
}

func TestEvalName(t *testing.T) {
	rdb := New(&RedisOption{
		Addrs:  []string{"192.168.123.100:6379"},
		Prefix: "aabbcc",
	})

	_, err := rdb.Eval(context.Background(), "return 1", []string{"a1", "a2", "a3"}, "b1", "b2", "b3").Result()
	assert.Nil(t, err, err)

	v, err := rdb.FunctionLoadReplace(context.Background(), "#!lua name=mylib\nredis.register_function('myfunc1', function() return 'hello world' end)").Result()
	t.Log(v, err)
}
