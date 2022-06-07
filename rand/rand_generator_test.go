package rand

import "testing"

func TestFastGenerate(t *testing.T) {
	g := NewFastRandGenerator()
	for i := 0; i < 100; i++ {
		t.Log(g.Int63())
	}
}

func BenchmarkGenerate(b *testing.B) {
	b.Run("fast", func(b *testing.B) {
		g1 := NewFastRandGenerator()
		for i := 0; i < b.N; i++ {
			g1.Int31()
		}
	})

	b.Run("normal", func(b *testing.B) {
		g1 := NewRandGenerator()
		for i := 0; i < b.N; i++ {
			g1.Int31()
		}
	})
}

func BenchmarkParallel(b *testing.B) {
	g1 := NewFastRandGenerator()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			g1.Int31()
		}
	})

	g2 := NewRandGenerator()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			g2.Int()
		}
	})
}
