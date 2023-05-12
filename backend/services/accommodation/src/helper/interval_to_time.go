package helper

import "time"

func GetDaysInInterval(startDate time.Time, endDate time.Time) []time.Time {
	var days []time.Time

	// Set the start date to the beginning of the day
	startDate = startDate.Truncate(24 * time.Hour)

	// Iterate over each day within the interval
	for currentDate := startDate; currentDate.Before(endDate); currentDate = currentDate.Add(24 * time.Hour) {
		// Append the current date to the array
		days = append(days, currentDate)
	}

	return days
}
