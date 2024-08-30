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
