package dates

import (
	"fmt"
	"time"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w]
}

type DateRange struct {
	BeginDate string
	EndDate   string
}

func FindDateRanges(startDate, endDate string, dayBegin, dayStart Weekday) ([]DateRange, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	if start.After(end) {
		return []DateRange{}, nil
	}

	var ranges []DateRange
	currentStart := start

	for currentStart.Before(end) || currentStart.Equal(end) {
		for Weekday(currentStart.Weekday()) != dayBegin {
			currentStart = currentStart.AddDate(0, 0, 1)
			if currentStart.After(end) {
				return ranges, nil
			}
		}

		currentEnd := currentStart
		for Weekday(currentEnd.Weekday()) != dayStart {
			currentEnd = currentEnd.AddDate(0, 0, 1)
		}

		if currentEnd.After(end) {
			currentEnd = end
		}

		ranges = append(ranges, DateRange{
			BeginDate: currentStart.Format("2006-01-02"),
			EndDate:   currentEnd.Format("2006-01-02"),
		})

		currentStart = currentEnd.AddDate(0, 0, 1)
	}

	return ranges, nil
}
