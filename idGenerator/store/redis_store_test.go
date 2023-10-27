package store_test

import (
	"testing"
	"time"

	"github.com/charlienet/go-mixed/idGenerator/store"
	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/tests"
)

func TestSmallSerail(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		s := store.NewRedisStore("sss", rdb)
		for i := 0; i < 5; i++ {
			t.Log(s.Assign(0, 9, 20))
		}
	}, redis.RedisOption{Addr: "192.168.123.50:6379", Password: "123456"})
}

func TestSmallAssign(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {

		s := store.NewRedisStore("sss", rdb)

		for i := 0; i < 10; i++ {
			t.Log(s.Assign(0, 9, 30))
		}

	}, redis.RedisOption{Addr: "192.168.123.50:6379", Password: "123456"})
}

func TestBigAssign(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {

		s := store.NewRedisStore("sss", rdb)

		for i := 0; i < 102; i++ {
			t.Log(s.Assign(0, 99, 10))
		}

	}, redis.RedisOption{Addr: "192.168.123.50:6379", Password: "123456"})
}

func TestRedisAssign(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {

		s := store.NewRedisStore("sss", rdb)

		for i := 0; i < 10; i++ {
			t.Log(s.Assign(21, 99, 30))
		}

	}, redis.RedisOption{Addr: "192.168.123.50:6379", Password: "123456"})
}

func TestFullRedisAssign(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {

		s := store.NewRedisStore("sss", rdb)

		for i := 0; i < 10; i++ {
			t.Log(s.Assign(0, 999, 99))
		}

	}, redis.RedisOption{Addr: "192.168.123.50:6379", Password: "123456"})
}

func TestUpdateMachineCode(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {

		for i := 0; i < 20; i++ {
			s := store.NewRedisStore("id", rdb)
			code, err := s.UpdateMachineCode(99)
			t.Log("获取到机器标识:", code, err)

			if err != nil {
				return
			}

			time.Sleep(time.Millisecond * 100)

			// s.Close()
		}

		time.Sleep(time.Second * 10)

	}, redis.RedisOption{Addr: "192.168.123.50:6379", Password: "123456", Prefix: "cacc"})

}

func TestUpdate(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		s := store.NewRedisStore("id", rdb)
		s.UpdateMachineCode(99)
		t.Log(s.MachineCode())

		s.UpdateMachineCode(99)
		t.Log(s.MachineCode())

		s2 := store.NewRedisStore("id", rdb)
		s2.UpdateMachineCode(99)
		t.Log(s2.MachineCode())

	}, redis.RedisOption{Addr: "192.168.123.50:6379", Password: "123456", Prefix: "cacc"})
}
