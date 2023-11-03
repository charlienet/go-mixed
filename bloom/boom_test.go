package bloom_test

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/charlienet/go-mixed/bloom"
	"github.com/charlienet/go-mixed/rand"
	"github.com/charlienet/go-mixed/redis"
	"github.com/charlienet/go-mixed/sys"
	"github.com/stretchr/testify/assert"
)

const ()

func TestBloom(t *testing.T) {
	b := bloom.New(1000, 0.03)

	for i := 0; i < 1000000; i++ {
		b.Add(context.Background(), strconv.Itoa(i))
	}

	v := "6943553521463296-1635402930"

	t.Log(b.ExistString(v))
	b.Add(context.Background(), v)
	t.Log(b.ExistString(v))

	isSet, err := b.ExistString(strconv.Itoa(9999))
	fmt.Println("过滤器中包含值:", isSet, err)

	isSet, err = b.ExistString("ss")
	fmt.Println("过滤器中未包含:", isSet, err)

	t.Log(sys.ShowMemUsage())
}

func TestOptimize(t *testing.T) {

	expectedInsertions := 1000000 // 期望存储数据量
	falseProbability := 0.00002   // 预期误差
	bits := uint(float64(-expectedInsertions) * math.Log(falseProbability) / (math.Log(2) * math.Log(2)))
	hashSize := uint(math.Round(float64(bits) / float64(expectedInsertions) * math.Log(2)))

	t.Log(bits)
	t.Log(hashSize)
}

func TestRedis(t *testing.T) {

	client := redis.New(&redis.RedisOption{
		Addrs:    []string{"192.168.2.222:6379"},
		Password: "123456",
	})

	bf := bloom.New(10000, 0.03, bloom.WithRedis(client, "bloom:test"))

	for i := 0; i < 100; i++ {
		bf.Add(context.Background(), strconv.Itoa(i))
	}

	for i := 0; i < 100; i++ {
		isSet, err := bf.ExistString(strconv.Itoa(i))
		if err != nil {
			t.Fatal(err)
		}

		if !isSet {
			t.Log(i, isSet)
		}
	}

	for i := 101; i < 200; i++ {
		isSet, err := bf.ExistString(strconv.Itoa(i))
		t.Log(isSet, err)
	}
}

func TestClear(t *testing.T) {
	bf := bloom.New(1000, 0.03)

	v := "abc"
	bf.Add(context.Background(), v)
	isSet, _ := bf.ExistString(v)
	assert.True(t, isSet)

	bf.Clear()
	isSet, _ = bf.ExistString(v)
	assert.False(t, isSet)
}

func TestParallel(t *testing.T) {
	f := bloom.New(1000, 0.03)

	for i := 0; i < 10000; i++ {
		v := rand.Hex.Generate(10)

		f.Add(context.Background(), v)
		isSet, _ := f.ExistString(v)

		assert.True(t, isSet)
	}
}

func BenchmarkFilter(b *testing.B) {
	f := bloom.New(1000, 0.03)

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {

			v := rand.Hex.Generate(10)
			f.Add(context.Background(), v)

			f.ExistString(v)

			// assert.True(b, f.Contains(v))

			// assert.True(b, f.Contains(v))
		}
	})

}
