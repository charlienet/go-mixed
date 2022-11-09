package calendar

import "time"

func Now() Calendar {
	return Create(time.Now())
}

func Today() Calendar {
	return Now().BeginningOfDay()
}

func Create(t time.Time) Calendar {
	return Calendar{
		Time:         t,
		weekStartsAt: WeekStartDay,
	}
}

func CreateFromTimestamp(timestamp int64) Calendar {
	return Create(time.Unix(timestamp, 0))
}

func CreateFromTimestampMilli(timestamp int64) Calendar {
	return Create(time.Unix(timestamp/1e3, (timestamp%1e3)*1e6))
}

func CreateFromTimestampMicro(timestamp int64) Calendar {
	return Create(time.Unix(timestamp/1e6, (timestamp%1e6)*1e3))
}

func CreateFromTimestampNano(timestamp int64) Calendar {
	return Create(time.Unix(timestamp/1e9, timestamp%1e9))
}

func create(year, month, day, hour, minute, second, nanosecond int) Calendar {
	return Calendar{
		Time: time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, time.Local),
	}
}
