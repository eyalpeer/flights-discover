package main

import (
	log "github.com/sirupsen/logrus"
	"main.go/analayze"
	"main.go/dates"
	"main.go/flights"
)

func main() {
	startDate := "2024-08-01"
	endDate := "2025-08-31"

	dayBegin := dates.Thursday
	dayEnd := dates.Sunday

	ranges, err := dates.FindDateRanges(startDate, endDate, dayBegin, dayEnd)
	if err != nil {
		log.Errorf("Failed to find date ranges: %v", err)
	}
	for _, r := range ranges {
		results, err := flights.SendRequest(r)
		if err != nil {
			log.Errorf("Failed to send request: %v", err)
		}
		bestPackages, err := analayze.AnalyzeFlightPackages(results, 400)
		if err != nil {
			log.Errorf("Failed to analyze flight packages: %v", err)
		}
		log.Info(flights.BeautifyResults(bestPackages, false))
	}
	log.Infof("Date ranges: %v", ranges)
}
