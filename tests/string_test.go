package tests

import (
	"bytes"
	"fmt"
	"net"
	"testing"
)

func BenchmarkStringSplice(b *testing.B) {
	userID := "aaaaa"
	orderID := "bbccc"

	b.Run("BenchmarkPlus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			logStr := "userid :" + userID + "; orderid:" + orderID
			_ = logStr
		}
	})

	b.Run("BenchmarkPrint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			logStr := fmt.Sprintf("userid: %v; orderid: %v", userID, orderID)
			_ = logStr
		}
	})

	b.Run("BenchmarkBytesBuffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var sb bytes.Buffer
			sb.WriteString("userid :")
			sb.WriteString(userID)
			sb.WriteString("; orderid:")
			sb.WriteString(orderID)

			logStr := sb.String()
			_ = logStr
		}
	})
}

func TestIPSegment(t *testing.T) {
	i, n, err := net.ParseCIDR("0.0.0.0/0")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(i, n)

	address := net.ParseIP("192.168.0.2")

	t.Log(n.Contains(address))
}
