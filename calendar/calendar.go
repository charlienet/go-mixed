package calendar

import (
	"time"
)

var WeekStartDay time.Weekday = time.Sunday

type Calendar struct {
	time.Time
	weekStartsAt time.Weekday
}

func BeginningOfMinute() Calendar {
	return Create(time.Now()).BeginningOfMinute()
}

func BeginningOfHour() Calendar {
	return Create(time.Now()).BeginningOfHour()
}

func BeginningOfDay() Calendar {
	return Create(time.Now()).BeginningOfDay()
}

func BeginningOfWeek() Calendar {
	return Create(time.Now()).BeginningOfWeek()
}

func BeginningOfMonth() Calendar {
	return Create(time.Now()).BeginningOfMonth()
}

func BeginningOfQuarter() Calendar {
	return Create(time.Now()).BeginningOfQuarter()
}

func BeginningOfYear() Calendar {
	return Create(time.Now()).BeginningOfYear()
}

func EndOfMinute() Calendar {
	return Create(time.Now()).EndOfMinute()
}

func EndOfHour() Calendar {
	return Create(time.Now()).EndOfHour()
}

func EndOfDay() Calendar {
	return Create(time.Now()).EndOfDay()
}

func EndOfWeek() Calendar {
	return Create(time.Now()).EndOfWeek()
}

func EndOfMonth() Calendar {
	return Create(time.Now()).EndOfMonth()
}

func EndOfQuarter() Calendar {
	return Create(time.Now()).EndOfQuarter()
}

func EndOfYear() Calendar {
	return Create(time.Now()).EndOfYear()
}

func (c Calendar) WeekStartsAt(day time.Weekday) Calendar {
	return Calendar{
		Time:         c.Time,
		weekStartsAt: day,
	}
}

func (c Calendar) BeginningOfMinute() Calendar {
	return Calendar{Time: c.Truncate(time.Minute)}

}

func (c Calendar) BeginningOfHour() Calendar {
	y, m, d := c.Date()
	return Calendar{
		Time:         time.Date(y, m, d, c.Hour(), 0, 0, 0, c.Location()),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) BeginningOfDay() Calendar {
	y, m, d := c.Date()

	return Calendar{
		Time: time.Date(y, m, d, 0, 0, 0, 0, c.Location()),
	}
}

func (c Calendar) BeginningOfWeek() Calendar {
	t := c.BeginningOfDay()
	weekday := int(t.Weekday())

	if c.weekStartsAt != time.Sunday {
		weekStartDayInt := int(c.weekStartsAt)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}

	return Calendar{
		Time:         t.AddDate(0, 0, -weekday),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) BeginningOfMonth() Calendar {
	y, m, _ := c.Date()

	return Calendar{
		Time:         time.Date(y, m, 1, 0, 0, 0, 0, c.Location()),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) BeginningOfQuarter() Calendar {
	month := c.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3

	return Calendar{
		Time:         month.AddDate(0, -offset, 0),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) BeginningOfYear() Calendar {
	y, _, _ := c.Date()

	return Calendar{
		Time:         time.Date(y, time.January, 1, 0, 0, 0, 0, c.Location()),
		weekStartsAt: c.weekStartsAt}
}

func (c Calendar) EndOfMinute() Calendar {
	n := c.BeginningOfMinute()

	return Calendar{
		Time:         n.Add(time.Minute - time.Nanosecond),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) EndOfHour() Calendar {
	n := c.BeginningOfHour()

	return Calendar{
		Time:         n.Add(time.Hour - time.Nanosecond),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) EndOfDay() Calendar {
	y, m, d := c.Date()

	return Calendar{
		Time:         time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location()),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) EndOfWeek() Calendar {
	n := c.BeginningOfWeek()

	return Calendar{
		Time:         n.AddDate(0, 0, 7).Add(-time.Nanosecond),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) EndOfMonth() Calendar {
	n := c.BeginningOfMonth()

	return Calendar{
		Time:         n.AddDate(0, 1, 0).Add(-time.Nanosecond),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) EndOfQuarter() Calendar {
	n := c.BeginningOfQuarter()

	return Calendar{
		Time:         n.AddDate(0, 3, 0).Add(-time.Nanosecond),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) EndOfYear() Calendar {
	n := c.BeginningOfYear()

	return Calendar{
		Time:         n.AddDate(1, 0, 0).Add(-time.Nanosecond),
		weekStartsAt: c.weekStartsAt,
	}
}

func (c Calendar) ToTime() time.Time {
	return c.Time
}
