package utils

import "time"

func FormatYYYY_MM_DD(dateStr string) string {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return ""
	}
	return date.Format("2006-01-02")
}
