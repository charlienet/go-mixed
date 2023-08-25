package calendar

import "time"

type ScheduledExecutor struct {
}

func NewScheduledExecutor() *ScheduledExecutor {
	return &ScheduledExecutor{}
}

func (e *ScheduledExecutor) Schedule(i any, duration time.Duration) {

}
