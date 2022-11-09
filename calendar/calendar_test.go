package calendar_test

import (
	"testing"
	"time"

	"github.com/charlienet/go-mixed/calendar"
	"github.com/stretchr/testify/assert"
)

var format = "2006-01-02 15:04:05.999999999"

func TestToday(t *testing.T) {
	t.Log(calendar.Today())
}

func TestBeginningOf(t *testing.T) {
	a := assert.New(t)

	n := time.Date(2022, 11, 9, 14, 28, 34, 123456789, time.UTC)
	a.Equal("2022-11-06 00:00:00", calendar.Create(n).BeginningOfWeek().String())
	a.Equal("2022-11-07 00:00:00", calendar.Create(n).WeekStartsAt(time.Monday).BeginningOfWeek().Format(format))
	a.Equal("2022-11-09 14:00:00", calendar.Create(n).BeginningOfHour().Format(format))
	a.Equal("2022-11-01 00:00:00", calendar.Create(n).BeginningOfMonth().Format(format))
	a.Equal("2022-10-01 00:00:00", calendar.Create(n).BeginningOfQuarter().Format(format))
}

func TestEndOf(t *testing.T) {
	a := assert.New(t)

	n := time.Date(2022, 11, 9, 14, 28, 34, 123456789, time.UTC)
	a.Equal("2022-11-09 14:28:59.999999999", calendar.Create(n).EndOfMinute().Format(format))
	a.Equal("2022-11-09 14:59:59.999999999", calendar.Create(n).EndOfHour().Format(format))
	a.Equal("2022-11-09 23:59:59.999999999", calendar.Create(n).EndOfDay().Format(format))
	a.Equal("2022-11-30 23:59:59.999999999", calendar.Create(n).EndOfMonth().Format(format))
	a.Equal("2022-12-31 23:59:59.999999999", calendar.Create(n).EndOfQuarter().Format(format))
	a.Equal("2022-12-31 23:59:59.999999999", calendar.Create(n).EndOfYear().Format(format))
}
