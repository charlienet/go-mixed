package currency

import (
	"testing"
)

func TestCentToDollar(t *testing.T) {
	cases := []struct {
		cent     int32
		excepted string
	}{
		{24040, "240.40"},
		{99999940, "999999.40"},
		{99999, "999.99"},
		{1, "0.01"},
		{99999901, "999999.01"},
		{100000099, "1000000.99"},
	}

	for _, c := range cases {
		result := CentToDollar(c.cent)
		if result != c.excepted {
			t.Fatalf("dollar to cent failed, dollar:%d execpted:%s result:%s", c.cent, c.excepted, result)
		}
	}
}

func TestDollarToCent(t *testing.T) {
	cases := []struct {
		dollar   string
		excepted int64
	}{
		{"240.40", 24040},
		{"999999.40", 99999940},
		{"999.99", 99999},
		{"0.01", 1},
		{"999999.01", 99999901},
		{"1000000.99", 100000099},
	}

	for _, c := range cases {
		result := DollarToCent(c.dollar)
		if result != c.excepted {
			t.Fatalf("dollar to cent failed, dollar:%s execpted:%d result:%d", c.dollar, c.excepted, result)
		}
	}
}

func TestYuanToFen(t *testing.T) {
	cases := []struct {
		dollar   string
		excepted int64
	}{
		{"240.40", 24040},
		{"999999.40", 99999940},
		{"999.99", 99999},
		{"0.01", 1},
		{"999999.01", 99999901},
		{"1000000.99", 100000099},
	}

	for _, c := range cases {
		result := YuanToFen(c.dollar)
		if result != c.excepted {
			t.Fatalf("dollar to cent failed, dollar:%s execpted:%d result:%d", c.dollar, c.excepted, result)
		}
	}
}

func TestFenToYuan(t *testing.T) {
	cases := []struct {
		cent     int
		excepted string
	}{
		{24040, "240.40"},
		{99999940, "999999.40"},
		{99999, "999.99"},
		{1, "0.01"},
		{99999901, "999999.01"},
		{100000099, "1000000.99"},
	}

	for _, c := range cases {
		result := FenToYuan(c.cent)
		if result != c.excepted {
			t.Fatalf("dollar to cent failed, dollar:%d execpted:%s result:%s", c.cent, c.excepted, result)
		}
	}
}
