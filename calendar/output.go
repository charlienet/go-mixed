package calendar

import "time"

// 布局模板常量
const (
	ANSICLayout              = time.ANSIC
	UnixDateLayout           = time.UnixDate
	RubyDateLayout           = time.RubyDate
	RFC822Layout             = time.RFC822
	RFC822ZLayout            = time.RFC822Z
	RFC850Layout             = time.RFC850
	RFC1123Layout            = time.RFC1123
	RFC1123ZLayout           = time.RFC1123Z
	RssLayout                = time.RFC1123Z
	KitchenLayout            = time.Kitchen
	RFC2822Layout            = time.RFC1123Z
	CookieLayout             = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC3339Layout            = "2006-01-02T15:04:05Z07:00"
	RFC3339MilliLayout       = "2006-01-02T15:04:05.999Z07:00"
	RFC3339MicroLayout       = "2006-01-02T15:04:05.999999Z07:00"
	RFC3339NanoLayout        = "2006-01-02T15:04:05.999999999Z07:00"
	ISO8601Layout            = "2006-01-02T15:04:05-07:00"
	ISO8601MilliLayout       = "2006-01-02T15:04:05.999-07:00"
	ISO8601MicroLayout       = "2006-01-02T15:04:05.999999-07:00"
	ISO8601NanoLayout        = "2006-01-02T15:04:05.999999999-07:00"
	RFC1036Layout            = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC7231Layout            = "Mon, 02 Jan 2006 15:04:05 GMT"
	DayDateTimeLayout        = "Mon, Jan 2, 2006 3:04 PM"
	DateTimeLayout           = "2006-01-02 15:04:05"
	DateTimeMilliLayout      = "2006-01-02 15:04:05.999"
	DateTimeMicroLayout      = "2006-01-02 15:04:05.999999"
	DateTimeNanoLayout       = "2006-01-02 15:04:05.999999999"
	ShortDateTimeLayout      = "20060102150405"
	ShortDateTimeMilliLayout = "20060102150405.999"
	ShortDateTimeMicroLayout = "20060102150405.999999"
	ShortDateTimeNanoLayout  = "20060102150405.999999999"
	DateLayout               = "2006-01-02"
	DateMilliLayout          = "2006-01-02.999"
	DateMicroLayout          = "2006-01-02.999999"
	DateNanoLayout           = "2006-01-02.999999999"
	ShortDateLayout          = "20060102"
	ShortDateMilliLayout     = "20060102.999"
	ShortDateMicroLayout     = "20060102.999999"
	ShortDateNanoLayout      = "20060102.999999999"
	TimeLayout               = "15:04:05"
	TimeMilliLayout          = "15:04:05.999"
	TimeMicroLayout          = "15:04:05.999999"
	TimeNanoLayout           = "15:04:05.999999999"
	ShortTimeLayout          = "150405"
	ShortTimeMilliLayout     = "150405.999"
	ShortTimeMicroLayout     = "150405.999999"
	ShortTimeNanoLayout      = "150405.999999999"
)

func (c Calendar) String() string {
	return c.ToDateTimeString()
}

func (c Calendar) ToDateTimeString() string {
	return c.Format(DateTimeLayout)
}

func (c Calendar) ToDateTimeInt() int {
	return c.ToShortDateInt()*1000000 + c.Hour()*10000 + c.Minute()*100 + c.Second()
}

func (c Calendar) ToShortDateInt() int {
	return c.Year()*10000 + int(c.Month())*100 + c.Day()
}

func (c Calendar) ToMonthInt() int {
	return c.Year()*100 + int(c.Month())
}
