package dates

import (
	"fmt"
	"strings"
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

type Range struct {
	BeginDate string
	EndDate   string
}

func FindDateRanges(startDate, endDate string, dayBegin, dayStart Weekday) ([]Range, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date: %v", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date: %v", err)
	}

	if start.After(end) {
		return []Range{}, nil
	}

	var ranges []Range
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

		ranges = append(ranges, Range{
			BeginDate: currentStart.Format("2006-01-02"),
			EndDate:   currentEnd.Format("2006-01-02"),
		})

		currentStart = currentEnd.AddDate(0, 0, 1)
	}

	return ranges, nil
}

func addOrSubtractDays(dateStr string, days int) (string, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}
	newDate := date.AddDate(0, 0, days)
	return newDate.Format("2006-01-02"), nil
}

func (dateRange Range) DateRangeToInput(expand int) (*Input, error) {
	departureStart, err := addOrSubtractDays(dateRange.BeginDate, -expand)
	if err != nil {
		return nil, err
	}
	departureEnd, err := addOrSubtractDays(dateRange.BeginDate, expand)
	if err != nil {
		return nil, err
	}
	arrivalStart, err := addOrSubtractDays(dateRange.EndDate, -expand)
	if err != nil {
		return nil, err
	}
	arrivalEnd, err := addOrSubtractDays(dateRange.EndDate, expand)
	if err != nil {
		return nil, err
	}
	departureStartSlice := strings.Split(departureStart, "-")
	departureEndSlice := strings.Split(departureEnd, "-")
	arrivalStartSlice := strings.Split(arrivalStart, "-")
	arrivalEndSlice := strings.Split(arrivalEnd, "-")
	return &Input{
		DepartureStartYear:  departureStartSlice[0],
		DepartureStartMonth: departureStartSlice[1],
		DepartureStartDay:   departureStartSlice[2],
		DepartureEndYear:    departureEndSlice[0],
		DepartureEndMonth:   departureEndSlice[1],
		DepartureEndDay:     departureEndSlice[2],
		ArriveStartYear:     arrivalStartSlice[0],
		ArriveStartMonth:    arrivalStartSlice[1],
		ArriveStartDay:      arrivalStartSlice[2],
		ArriveEndYear:       arrivalEndSlice[0],
		ArriveEndMonth:      arrivalEndSlice[1],
		ArriveEndDay:        arrivalEndSlice[2],
	}, nil
}
