package appUtil

import "time"

func GetLastMonth() (time.Month, int) {
	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()

	lastMonth := currentMonth - 1
	lastYear := currentYear
	if lastMonth == 0 {
		lastMonth = 12
		lastYear--
	}

	return lastMonth, lastYear
}
