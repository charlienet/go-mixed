package rand

import "testing"

func TestSecureInt63n(t *testing.T) {
	r := NewSecureRandGenerator()
	for i := 0; i < 100; i++ {
		t.Log(r.Int63n(100000))
	}
}

func TestSecureInt63(t *testing.T) {
	r := NewSecureRandGenerator()
	for i := 0; i < 100; i++ {
		t.Log(r.Int63())
	}
}
