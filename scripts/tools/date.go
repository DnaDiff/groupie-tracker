package tools

import (
	"strings"
	"time"
)

// Convert the dates to the correct format
// EX: "2019-01-01" -> "January 1, 2019"
func Date(date string) string {
	// 02, 01, 2006 -> 02-01-2006
	date = strings.ReplaceAll(date, ", ", "-")
	// Parse the date string using the "DD-MM-YYYY" layout
	t, err := time.Parse("02-01-2006", date)
	if err != nil {
		return date
	}
	// Format the time using the "Month D, YYYY" layout
	return t.Format("January 2, 2006")
}
