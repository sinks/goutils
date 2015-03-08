package week

import "time"

// Start returns a date at the start of the week
func Start(t time.Time, start time.Weekday) time.Time {
	shift := ShiftDays(t.Weekday(), start)
	return t.AddDate(0, 0, shift)
}

// ShiftDays calculates how many days between c back to start
func ShiftDays(c time.Weekday, start time.Weekday) int {
	shift := int(start-c) % 7
	if shift > 0 {
		shift = shift - 7
	}
	return shift
}
