package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	pwd := "123456"

	h1, _ := GenerateFromPassword([]byte(pwd))
	t.Log(ComparePassword(h1, []byte(pwd)))

	for i := 0; i < 100; i++ {
		h, err := GenerateFromPassword([]byte(pwd))
		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, ComparePassword(h, []byte(pwd)))
	}
}
