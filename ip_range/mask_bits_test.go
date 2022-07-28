package iprange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskToBits(t *testing.T) {

	masks := []struct {
		mask   string
		expect int
	}{
		{"255.255.255.0", 24},
		{"255.255.248.0", 21},
		{"255.255.192.0", 18},
		{"255.255.255.192", 26},
	}

	for _, m := range masks {
		bits := MaskToBits(m.mask)
		assert.Equal(t, m.expect, bits, fmt.Sprintf("IP:%s 掩码位数错误。", m.mask))
	}
}
