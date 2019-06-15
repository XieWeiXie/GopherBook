package activity

import "time"

func toTime(value string) time.Time {
	// format: 2006-01-02 15:04:05
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	return t
}
