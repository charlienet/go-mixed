package bloom_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/charlienet/go-mixed/bloom"
	"github.com/charlienet/go-mixed/rand"
	"github.com/charlienet/go-mixed/sys"
	"github.com/stretchr/testify/assert"
)

func TestBloom(t *testing.T) {
	b := bloom.NewBloomFilter()

	for i := 0; i < 1000000; i++ {
		b.Add(strconv.Itoa(i))
	}

	v := "6943553521463296-1635402930"

	t.Log(b.Contains(v))
	b.Add(v)
	t.Log(b.Contains(v))

	fmt.Println("过滤器中包含值:", b.Contains(strconv.Itoa(9999)))
	fmt.Println("过滤器中未包含:", b.Contains("ss"))

	t.Log(sys.ShowMemUsage())
}

func TestSize(t *testing.T) {
	bloom.NewBloomFilter(bloom.WithSize(1 << 2))
}

func TestClear(t *testing.T) {
	bf := bloom.NewBloomFilter()

	v := "abc"
	bf.Add(v)
	assert.True(t, bf.Contains(v))

	bf.Clear()
	assert.False(t, bf.Contains(v))
}

func TestParallel(t *testing.T) {
	f := bloom.NewBloomFilter()

	for i := 0; i < 10000; i++ {
		v := rand.Hex.Generate(10)

		f.Add(v)
		assert.True(t, f.Contains(v))
	}
}

func BenchmarkFilter(b *testing.B) {
	f := bloom.NewBloomFilter()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			v := rand.Hex.Generate(10)
			f.Add(v)

			f.Contains(v)

			// assert.True(b, f.Contains(v))

			// assert.True(b, f.Contains(v))
		}
	})

}
