package internal

import "time"

// NOTE: in golang the layout to use with Time.Format() relies
// on "2006-01-02T15:04:05.000Z"
const (
	Layout_yyyyMMdd = "2006-01-02"
	Layout_hhmm     = "15:04"
)

func toDateString(sourceTime *time.Time) string {
	return sourceTime.Local().Format(Layout_yyyyMMdd)
}

func toShortTime(sourceTime *time.Time) string {
	return sourceTime.Local().Format(Layout_hhmm)
}

func toTodayZeroOClock(sourceTime *time.Time) time.Time {
	return time.Date(sourceTime.Local().Year(), sourceTime.Local().Month(), sourceTime.Local().Day(), 0, 0, 0, 0, time.Local)
}
