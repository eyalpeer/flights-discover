package main

import (
	log "github.com/sirupsen/logrus"
	"main.go/dates"
)

func main() {
	startDate := "2023-06-01"
	endDate := "2023-06-30"
	dayBegin := dates.Thursday
	dayEnd := dates.Sunday

	ranges, err := dates.FindDateRanges(startDate, endDate, dayBegin, dayEnd)
	if err != nil {
		log.Errorf("Failed to find date ranges: %v", err)
	}
	log.Infof("Date ranges: %v", ranges)
}
