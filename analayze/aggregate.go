package analayze

import (
	"fmt"
	"main.go/flights"
	"time"
)

const minimumDuration = 2

func AnalyzeFlightPackages(packages *flights.FlightPackage) (map[string]flights.Route, error) {
	cheapestPackages := make(map[string]flights.Route)
	dateRanges := make(map[string][]time.Time)

	for _, pkg := range (*packages).Packages {
		for _, route := range pkg.Route {
			departureDate := time.Unix(route.ATime, 0)
			arrivalDate := time.Unix(route.DrTime, 0)
			if arrivalDate.Sub(departureDate) < minimumDuration {
				continue
			}
			key := fmt.Sprintf("%s_%s_%s", route.FlyTo, departureDate.Format("2006-01-02"), arrivalDate.Format("2006-01-02"))
			if existingRoute, exists := cheapestPackages[key]; !exists || route.Price < existingRoute.Price {
				cheapestPackages[key] = route
			}
			dateRanges[route.FlyTo] = append(dateRanges[route.FlyTo], departureDate, arrivalDate)
		}
	}
	return cheapestPackages, nil
}
