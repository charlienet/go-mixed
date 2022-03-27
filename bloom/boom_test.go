package bloom_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/charlienet/go-mixed/bloom"
)

func TestBloom(t *testing.T) {
	b := bloom.NewBloomFilter()

	for i := 0; i < 1000000; i++ {
		b.Add(strconv.Itoa(i))
	}

	fmt.Println(b.Contains(strconv.Itoa(9999)))
	fmt.Println(b.Contains("ss"))
}
