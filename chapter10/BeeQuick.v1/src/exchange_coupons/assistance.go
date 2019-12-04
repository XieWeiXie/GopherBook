package exchange_coupons

import (
	"fmt"
	"time"
)

func toTime(value string) (time.Time, error) {
	var (
		timeValue time.Time
		err       error
	)
	if len(value) == 10 {
		value = fmt.Sprintf("%s 00:00:00", value)
	}
	v, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	if err != nil {
		return timeValue, err
	}
	return v, nil
}
