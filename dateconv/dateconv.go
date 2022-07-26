package dateconv

import (
	"time"
)

const (
	layoutDate        = "2006-01-02"
	layoutTime        = "2006-01-02 15:04:05"
	layoutTimeMilli   = "2006-01-02 15:04:05.000"
	layoutChineseDate = "2006年01月02日"
	layoutChineseTime = "2006年01月02日 15:04:05"
)

func Today() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// 日期转换为整数（如:20211222）
func DateToInt(date time.Time) int {
	return date.Year()*10000 + int(date.Month())*100 + date.Day()
}

// 日期转换为字符串
func DateToString(date *time.Time) string { return formatTime(date, layoutDate) }

// 时间转换为字符串
func TimeToString(date *time.Time) string { return formatTime(date, layoutTime) }

// 日期转换为中文
func DateToChinese(t *time.Time) string { return formatTime(t, layoutChineseDate) }

// 时间转换为中文
func TimeToChinese(t *time.Time) string { return formatTime(t, layoutChineseTime) }

func formatTime(t *time.Time, f string) string {
	if t == nil || t.IsZero() {
		return ""
	}

	return t.Format(f)
}
