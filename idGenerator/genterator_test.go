package idgenerator_test

import (
	"sync"
	"testing"

	idgenerator "github.com/charlienet/go-mixed/idGenerator"
	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/sets"
	"github.com/charlienet/go-mixed/tests"
)

var redisOption = redis.ReidsOption{Addr: "192.168.123.50:6379", Password: "123456"}

func TestGenerator(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		generator, err := idgenerator.New(
			idgenerator.WithDecimalFormater(idgenerator.YYYYMMDDHHmmss, 1, 1),
			idgenerator.WithRedis("idgen_test", rdb))
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 20; i++ {
			t.Log(generator.Next())
		}
	})
}

func TestDecimalGenerator(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		generator, err := idgenerator.New(
			idgenerator.WithDecimalFormater(idgenerator.YYYYMMDDHHmmss, 3, 1),
			idgenerator.WithRedis("idgen_test", rdb))
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 200; i++ {
			t.Log(generator.Next())
		}

	}, redis.ReidsOption{Addr: "192.168.123.50:6379", Password: "123456"})
}

func TestDecimalMonth(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		generator, err := idgenerator.New(
			idgenerator.WithDecimalFormater(idgenerator.YYYYMMDD, 2, 1),
			idgenerator.WithRedis("idgen_test", rdb))
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 105; i++ {
			t.Log(generator.Next())
		}

	}, redis.ReidsOption{Addr: "192.168.123.50:6379", Password: "123456"})
}

func TestParallelCreate(t *testing.T) {
	tests.RunOnRedis(t, func(rdb redis.Client) {
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()

			g1, err := idgenerator.New(
				idgenerator.WithDecimalFormater(idgenerator.YYYYMMDDHHmmss, 3, 1),
				idgenerator.WithRedis("idgen_testcccc", rdb))
			if err != nil {
				panic(err)
			}

			_ = g1.Next().Int64()
		}()

		go func() {
			defer wg.Done()

			g2, err := idgenerator.New(
				idgenerator.WithDecimalFormater(idgenerator.YYYYMMDDHHmmss, 3, 1),
				idgenerator.WithRedis("idgen_testcccc", rdb))
			if err != nil {
				panic(err)
			}

			_ = g2.Next().Int64()
		}()

		wg.Wait()

	}, redisOption)
}

func TestParallel(t *testing.T) {
	set := sets.NewHashSet[int64]().Sync()
	opt := redis.ReidsOption{Addr: "192.168.123.50:6379", Password: "123456"}

	_ = set
	f := func() {
		tests.RunOnRedis(t, func(rdb redis.Client) {
			generator, err := idgenerator.New(
				idgenerator.WithDecimalFormater(idgenerator.YYYYMMDDHHmmss, 3, 1),
				idgenerator.WithRedis("idgen_testcccc", rdb))
			if err != nil {
				t.Fatal(err)
			}
			defer generator.Close()

			generator.Next()
			for i := 0; i < 50000; i++ {
				id := generator.Next().Int64()
				if set.Contains(id) {
					panic("生成重复")
				}
				set.Add(id)
			}

		}, opt)
	}

	var wg sync.WaitGroup
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}

	wg.Wait()
}

func BenchmarkGenerator(b *testing.B) {
	tests.RunOnRedis(b, func(rdb redis.Client) {
		b.Run("bbb", func(b *testing.B) {
			generator, err := idgenerator.New(
				idgenerator.WithDecimalFormater(idgenerator.YYYYMMDDHHmmss, 3, 1),
				idgenerator.WithRedis("idgen_test", rdb))
			if err != nil {
				b.Fatal(err)
			}

			for i := 0; i < 999; i++ {
				generator.Next()
			}

		})
	}, redis.ReidsOption{Addr: "192.168.123.50:6379", Password: "123456"})
}
