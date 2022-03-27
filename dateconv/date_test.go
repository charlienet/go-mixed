package dateconv_test

import (
	"testing"

	"github.com/charlienet/go-mixed/dateconv"
)

func TestToday(t *testing.T) {
	today := dateconv.Today()
	t.Log(dateconv.TimeToString(&today))
}
