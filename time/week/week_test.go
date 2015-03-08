package week

import (
	"testing"
	"time"
)

func TestShiftDays(t *testing.T) {
	type TestData struct {
		Start    time.Weekday
		Test     time.Weekday
		Expected int
	}
	data := []TestData{
		{time.Sunday, time.Monday, -1},
		{time.Sunday, time.Tuesday, -2},
		{time.Sunday, time.Wednesday, -3},
		{time.Sunday, time.Thursday, -4},
		{time.Sunday, time.Friday, -5},
		{time.Sunday, time.Saturday, -6},
		{time.Sunday, time.Sunday, 0},
		{time.Monday, time.Monday, 0},
		{time.Monday, time.Tuesday, -1},
		{time.Monday, time.Wednesday, -2},
		{time.Monday, time.Thursday, -3},
		{time.Monday, time.Friday, -4},
		{time.Monday, time.Saturday, -5},
		{time.Monday, time.Sunday, -6},
		{time.Tuesday, time.Monday, -6},
		{time.Tuesday, time.Tuesday, 0},
		{time.Tuesday, time.Wednesday, -1},
		{time.Tuesday, time.Thursday, -2},
		{time.Tuesday, time.Friday, -3},
		{time.Tuesday, time.Saturday, -4},
		{time.Tuesday, time.Sunday, -5},
		{time.Wednesday, time.Monday, -5},
		{time.Wednesday, time.Tuesday, -6},
		{time.Wednesday, time.Wednesday, 0},
		{time.Wednesday, time.Thursday, -1},
		{time.Wednesday, time.Friday, -2},
		{time.Wednesday, time.Saturday, -3},
		{time.Wednesday, time.Sunday, -4},
		{time.Thursday, time.Monday, -4},
		{time.Thursday, time.Tuesday, -5},
		{time.Thursday, time.Wednesday, -6},
		{time.Thursday, time.Thursday, 0},
		{time.Thursday, time.Friday, -1},
		{time.Thursday, time.Saturday, -2},
		{time.Thursday, time.Sunday, -3},
		{time.Saturday, time.Monday, -2},
		{time.Saturday, time.Tuesday, -3},
		{time.Saturday, time.Wednesday, -4},
		{time.Saturday, time.Thursday, -5},
		{time.Saturday, time.Friday, -6},
		{time.Saturday, time.Saturday, 0},
		{time.Saturday, time.Sunday, -1},
	}
	for _, data := range data {
		result := ShiftDays(data.Test, data.Start)
		if result != data.Expected {
			t.Errorf(
				"ShiftDays(%s, %s) == %d, expected: %d",
				data.Test,
				data.Start,
				result,
				data.Expected,
			)
		}
	}

}

func TestStart(t *testing.T) {
	type TestData struct {
		Start    time.Weekday
		TestDate time.Time
		Expected time.Time
	}
	data := []TestData{
		{
			time.Monday,
			time.Date(2015, time.March, 6, 0, 0, 0, 1, time.UTC), // Friday
			time.Date(2015, time.March, 2, 0, 0, 0, 1, time.UTC), // Monday
		},
		{
			time.Monday,
			time.Date(2015, time.March, 2, 0, 0, 0, 1, time.UTC), // Monday
			time.Date(2015, time.March, 2, 0, 0, 0, 1, time.UTC), // Monday
		},
		{
			time.Sunday,
			time.Date(2015, time.March, 6, 0, 0, 0, 1, time.UTC), // Friday
			time.Date(2015, time.March, 1, 0, 0, 0, 1, time.UTC), // Sunday
		},
	}
	for _, data := range data {
		result := Start(data.TestDate, data.Start)
		if result != data.Expected {
			t.Errorf(
				"WeekStart(%s, %s) == %s, expected: %s",
				data.TestDate,
				data.Start,
				result,
				data.Expected,
			)
		}
	}
}
