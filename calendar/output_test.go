package calendar_test

import (
	"testing"
	"time"

	"github.com/charlienet/go-mixed/calendar"
	"github.com/stretchr/testify/assert"
)

func TestDayInt(t *testing.T) {
	assert := assert.New(t)

	n := time.Date(2022, 11, 9, 14, 28, 34, 123456789, time.UTC)

	assert.Equal(20221109, calendar.Create(n).ToShortDateInt())
	assert.Equal(202211, calendar.Create(n).ToMonthInt())
	assert.Equal(20221109142834, calendar.Create(n).ToDateTimeInt())
}
