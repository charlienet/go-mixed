package iprange

import (
	"net/netip"
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

func TestSinglePrefix(t *testing.T) {
	r, err := NewRange("192.168.2.100/32")
	if err != nil {
		t.Fatal(err)
	}

	assert.False(t, r.Contains("192.168.2.56"))
	assert.True(t, r.Contains("192.168.2.100"))
	assert.False(t, r.Contains("192.168.2.130"))
}

func TestAllIp(t *testing.T) {
	r, err := NewRange("0.0.0.0/0")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("192.168.2.100"))
	assert.True(t, r.Contains("192.3.2.100"))
	assert.True(t, r.Contains("192.65.2.100"))
	assert.True(t, r.Contains("172.168.2.100"))
	assert.True(t, r.Contains("8.8.8.8"))
	assert.True(t, r.Contains("114.114.114.114"))
}

func TestPrefix(t *testing.T) {
	r, err := NewRange("192.168.2.0/24")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("192.168.2.12"))
	assert.True(t, r.Contains("192.168.2.162"))
	assert.False(t, r.Contains("192.168.3.162"))
}

func TestPrefix2(t *testing.T) {
	r, err := NewRange("192.168.15.0/21")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("192.168.8.10"))
	assert.True(t, r.Contains("192.168.14.162"))
	assert.False(t, r.Contains("192.168.3.162"))
	assert.False(t, r.Contains("192.168.2.162"))
}

func TestDotMask(t *testing.T) {
	r, err := NewRange("192.168.15.0/255.255.248.0")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("192.168.8.10"))
	assert.True(t, r.Contains("192.168.14.162"))
	assert.False(t, r.Contains("192.168.3.162"))
	assert.False(t, r.Contains("192.168.2.162"))
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

func TestLocalhost(t *testing.T) {
	r, err := NewRange("::1")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, r.Contains("::1"))
}

func TestNetIP(t *testing.T) {
	addr, err := netip.ParseAddr("192.168.2.10")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(netip.MustParseAddr("192.168.2.4").Compare(addr))
	t.Log(netip.MustParseAddr("192.168.2.10").Compare(addr))
	t.Log(netip.MustParseAddr("192.168.2.11").Compare(addr))

	prefix := netip.MustParsePrefix("192.168.2.0/24")

	t.Log(prefix.Contains(netip.MustParseAddr("192.168.2.53")))
	t.Log(prefix.Contains(netip.MustParseAddr("192.168.3.53")))
}
