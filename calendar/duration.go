package calendar

import "time"

func ParseDuration(s string) (time.Duration, error) {
	if len(s) == 0 {
		return time.Duration(0), nil
	}

	return time.ParseDuration(s)
}
