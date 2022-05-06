package rand_test

import (
	"bytes"
	"testing"
	"time"

	mrnd "math/rand"

	"github.com/charlienet/go-mixed/rand"
)

func TestRandString(t *testing.T) {
	t.Log(rand.AllChars.Generate(20))

	// b, err := rand.RandBytes(32)
	// t.Log(err)
	// t.Log(hex.EncodeToString(b))
}

func TestRandHex(t *testing.T) {
	h := rand.Hex.Generate(8)
	t.Log(h)
}

func TestRandMax(t *testing.T) {
	mrnd.Seed(time.Now().UnixNano())
}

func TestCryptRand(t *testing.T) {
	t.Log(rand.CryptoIntn(56545))
}

func TestGenericsInterger(t *testing.T) {
	var max int32 = 55

	for i := 0; i < 1000; i++ {
		if rand.Intn(max) >= max {
			t.Fatal("生成的值大于最大值:", max)
		}
	}
}

func TestRange(t *testing.T) {
	min := 20
	max := 200000
	for i := 0; i < 100000000; i++ {
		n := rand.IntRange(min, max)
		if n < min {
			t.Fatal("生成的值小于最小值:", min, n)
		}

		if n >= max {
			t.Fatal("生成的值大于最大值:", min, n)
		}
	}
}

func TestGenerator(t *testing.T) {
	g := rand.NewRandGenerator()

	for i := 0; i < 100; i++ {
		t.Log(g.Int63())
	}
}

func TestMutiGenerator(t *testing.T) {
	set := make(map[int64]struct{}, 1000)

	for i := 0; i < 1000; i++ {
		r := rand.NewRandGenerator().Int63()
		if _, ok := set[r]; ok {
			t.Fatal("生成的随机数重复")
		}

		set[r] = struct{}{}
	}
}

func BenchmarkParallel(b *testing.B) {
	rand.Hex.Generate(16)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Hex.Generate(16)
		}
	})
}

func BenchmarkNoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.AllChars.Generate(16)
	}
}

func BenchmarkHexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.RandBytes(16)
	}
}

func BenchmarkHexParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.RandBytes(16)
		}
	})
}

func BenchmarkGenerator(b *testing.B) {
	r := rand.NewRandGenerator()
	for i := 0; i < b.N; i++ {
		r.Int63()
	}
}

func BenchmarkParallelGenerator(b *testing.B) {
	r := rand.NewRandGenerator()
	for i:=0; i<10; i++{
		go r.Int63()
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Int63()
		}
	})
}

func BenchmarkRandString(b *testing.B) {
	b.Log(rand.Hex.Generate(16))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rand.Hex.Generate(16)
	}
}

func BenchmarkString(b *testing.B) {
	elems := []byte("abcdefghijk")

	b.Run("1", func(b *testing.B) {
		a := []byte{}
		for i := 0; i < b.N; i++ {
			for _, elem := range elems {
				a = append(a, elem)
			}
		}
	})

	b.Run("2", func(b *testing.B) {
		a := make([]byte, len(elems))
		for i := 0; i < b.N; i++ {
			for _, elem := range elems {
				a = append(a, elem)
			}
		}
	})

	b.Run("3", func(b *testing.B) {
		a := make([]byte, len(elems))
		for i := 0; i < b.N; i++ {
			a = append(a, elems...)
		}
	})
}

func BenchmarkConcatString(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}

	b.Run("add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ret := ""
			for _, elem := range elems {
				ret += elem
			}
		}
	})

	b.Run("buffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			for _, elem := range elems {
				buf.WriteString(elem)
			}
		}
	})
}

func BenchmarkRand(b *testing.B) {
	b.Run("math/rand", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rand.Intn(100)
		}
	})

	b.Run("crypto/rand", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			rand.CryptoIntn(100)
		}
	})

}
