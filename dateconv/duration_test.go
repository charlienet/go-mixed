package dateconv

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	t.Log(ParseDuration(""))
	t.Log(ParseDuration("abc"))
}

func TestMonth(t *testing.T) {
	month := time.Now()

	offset := (int(month.Month()) - 3)
	t.Log(offset)

	month = month.AddDate(0, -3, 1)
	t.Log(month)

	tt := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
	t.Log(tt)
}
