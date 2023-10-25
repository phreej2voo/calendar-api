package tool

import "time"

func MonthDuration(startTime, endTime time.Time) (month, day int) {
	duration := endTime.Sub(startTime)
	hours := duration.Hours()
	days := int(hours/24) + 1
	month = int(days / 30)
	day = days - month*30
	return
}
