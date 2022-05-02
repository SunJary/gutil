package str

import (
	"fmt"
	"time"
)

func GetDatePath() string {
	now := time.Now()
	return fmt.Sprintf("%d%02d/%02d", now.Year(), now.Month(), now.Day())
}

func DurationToString(duration time.Duration) (durationString string) {
	if duration < time.Minute {
		return fmt.Sprintf("%d分钟", int(duration.Minutes()))
	}
	hours := int(duration.Hours())
	d := hours / 24
	m := d / 30
	y := d / 365

	day := 24 * time.Hour
	if duration < day {
		return fmt.Sprintf("%d小时", hours)
	}

	month := 30 * day
	if duration < month {
		return fmt.Sprintf("%d天", d)
	}

	year := 365 * day
	if duration < year {
		monthStr := fmt.Sprintf("%d月", m)
		remindDays := d - 30*m
		if remindDays > 0 {
			return fmt.Sprintf(monthStr+"%d天", remindDays)
		}
		return monthStr
	}

	yearStr := fmt.Sprintf("%d年", y)

	if remindMonths := (d - 365*y) / 30; remindMonths > 0 {
		return fmt.Sprintf(yearStr+"%d月", remindMonths)
	} else if remindDays := (d - 365*y); remindDays > 0 {
		return fmt.Sprintf(yearStr+"%d天", remindDays)
	}
	return yearStr

}
