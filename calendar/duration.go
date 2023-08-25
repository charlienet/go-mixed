package calendar

import "time"

func ParseDuration(s string) (time.Duration, error) {
	if len(s) == 0 {
		return time.Duration(0), nil
	}

	return time.ParseDuration(s)
}

func ParseDurationDefault(s string, d time.Duration) time.Duration {
	if len(s) == 0 {
		return d
	}

	ret, err := time.ParseDuration(s)
	if err != nil {
		return d
	}

	return ret
}
