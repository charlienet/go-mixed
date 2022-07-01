package iprange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleErrorIP(t *testing.T) {
	values := []string{
		"192.168.01",
		"::",
	}

	for _, v := range values {
		r, err := NewRange(v)

		t.Log(r, err)
	}
}

func TestSingleIp(t *testing.T) {
	r, err := NewRange("192.168.0.1")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("192.168.0.1"))
	assert.False(t, r.Contains("192.168.0.123"))
}

func TestCIDR(t *testing.T) {
	r, err := NewRange("192.168.2.0/24")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("192.168.2.12"))
	assert.True(t, r.Contains("192.168.2.162"))
	assert.False(t, r.Contains("192.168.3.162"))
}

func TestRange(t *testing.T) {
	r, err := NewRange("192.168.2.20-192.168.2.30")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("192.168.2.20"))
	assert.True(t, r.Contains("192.168.2.21"))
	assert.True(t, r.Contains("192.168.2.30"))

	assert.False(t, r.Contains("192.168.2.10"))
	assert.False(t, r.Contains("192.168.2.31"))

}
