package validate

import "time"

func IsValidPeriod(period string) (isValid bool) {
	_, err := time.Parse("2006-01", period)
	if err != nil {
		return
	}
	isValid = true
	return
}

func IsValidDate(date string) (isValid bool) {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return
	}
	isValid = true
	return
}
