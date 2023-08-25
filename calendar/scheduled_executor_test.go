package calendar_test

import (
	"testing"
	"time"

	"github.com/charlienet/go-mixed/calendar"
)

func TestExecutor(t *testing.T) {
	executor := calendar.NewScheduledExecutor()

	executor.Schedule(nil, time.Minute)
}
