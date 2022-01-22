package util

import (
	"fmt"
	"time"
)

func FormatDate(date time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", date.Year(), date.Month(), date.Day())
}

func PlainDate(date time.Time) string {
	return fmt.Sprintf("%d %s %d", date.Day(), date.Month().String(), date.Year())
}

func PrettyDate(date time.Time) string {
	now := time.Now()
	day := date.YearDay()
	today := now.YearDay()

	if date.Year() == now.Year() && day <= today {
		// Friendly date formats
		if day == today {
			return "today"
		} else if day+1 == today {
			return "yesterday"
		} else if day+7 > today {
			return fmt.Sprintf("%d days ago", today-day)
		} else if day+14 > today {
			return "a week ago"
		} else if day+35 > today {
			return fmt.Sprintf("%d weeks ago", (today-day)/7)
		}

		// Posted this year
		return fmt.Sprintf("on %d %s", date.Day(), date.Month().String())
	}

	// Posted in a past year, or in the future
	return fmt.Sprintf("on %s", PlainDate(date))
}
