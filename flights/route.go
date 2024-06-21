package flights

import (
	"bytes"
	"fmt"
	"github.com/leonm1/airports-go"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Route struct {
	Route                []Segment   `json:"route"`
	Conversion           interface{} `json:"conversion"`
	DeepLink             string      `json:"deep_link"`
	Baggage              Baggage     `json:"baggage"`
	Price                float64     `json:"price"`
	FlyDuration          string      `json:"fly_duration"`
	FlyDurationSecs      float64     `json:"fly_duration_secs"`
	FlyFrom              string      `json:"flyFrom"`
	FlyFromCity          string
	FlyFromAirport       string
	FlyTo                string `json:"flyTo"`
	FlyToCity            string
	FlyToAirport         string
	DTime                int64       `json:"dTime"`
	ATime                int64       `json:"aTime"`
	ReturnDuration       string      `json:"return_duration"`
	ReturnDurationSecs   float64     `json:"return_duration_secs"`
	NightsInDest         interface{} `json:"nightsInDest"`
	CityCodeFrom         string      `json:"cityCodeFrom"`
	CityCodeTo           string      `json:"cityCodeTo"`
	OperatingCarrier     string      `json:"operating_carrier"`
	OperatingCarrierName string
	Sc                   string   `json:"sc"`
	ArrivalPos           int      `json:"arrivalpos"`
	TripSpec             []string `json:"tripspec"`
	Airports             string   `json:"airports"`
	EndAirports          string   `json:"endairports"`
	Airlines             string   `json:"airlines"`
	MaxStops             int      `json:"maxstops"`
	Co2                  int      `json:"co2"`
	FlyDurationMin       float64  `json:"fly_duration_min"`
	SegDurMax            float64  `json:"segdurmax"`
	DrTime               int64    `json:"drTime"`
	ArTime               int64    `json:"arTime"`
	ReturnDurationMin    float64  `json:"return_duration_min"`
}

func (route *Route) Enrich() {
	airport, err := airports.LookupIATA(route.FlyFrom)
	if err != nil {
		log.Warnf("Failed to lookup airport: %v", err)
		route.FlyFromCity = route.FlyFrom
		route.FlyFromAirport = route.FlyFrom
	} else {
		route.FlyFromCity = airport.City
		route.FlyFromAirport = airport.Name
	}
	airport, err = airports.LookupIATA(route.FlyTo)
	if err != nil {
		log.Warnf("Failed to lookup airport: %v", err)
		route.FlyToCity = route.FlyTo
		route.FlyToAirport = route.FlyTo
	} else {
		route.FlyToCity = airport.City
		route.FlyToAirport = airport.Name
	}
	airline, err := FindAirlineName(route.OperatingCarrier)
	if err != nil {
		log.Warnf("Failed to lookup airline: %v", err)
		route.OperatingCarrierName = route.OperatingCarrier
	} else {
		route.OperatingCarrierName = airline
	}
}

func (route *Route) String() string {
	route.Enrich()

	var depTime, arrTime, retDepTime, retArrTime string
	if departureTime := time.Unix(route.DTime, 0); !departureTime.IsZero() {
		depTime = departureTime.Format("2006-01-02 15:04")
	}
	if arrivalTime := time.Unix(route.ATime, 0); !arrivalTime.IsZero() {
		arrTime = arrivalTime.Format("2006-01-02 15:04")
	}
	if returnDepartureTime := time.Unix(route.DrTime, 0); !returnDepartureTime.IsZero() {
		retDepTime = returnDepartureTime.Format("2006-01-02 15:04")
	}
	if returnArrivalTime := time.Unix(route.ArTime, 0); !returnArrivalTime.IsZero() {
		retArrTime = returnArrivalTime.Format("2006-01-02 15:04")
	}

	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%d",
		depTime, arrTime, retDepTime, retArrTime,
		route.FlyToCity, route.OperatingCarrierName, int(math.Round(route.Price)),
	)
}

func BeautifyResults(results map[string]Route, links bool) string {
	if len(results) == 0 {
		return "No results found."
	}
	sortSlice := sortMap(results)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Departure\tArrival\tReturn Dep.\tReturn Arr.\tDestination\tCarrier\tPrice")
	fmt.Fprintln(w, "---------\t-------\t----------\t----------\t-----------\t-------\t-----")

	for _, route := range sortSlice {
		routeDetails := route.String()
		fmt.Fprintln(w, routeDetails)
	}

	w.Flush()
	if links {
		var linksBuffer bytes.Buffer
		for _, route := range sortSlice {
			if route.DeepLink != "" {
				linksBuffer.WriteString(route.DeepLink)
				linksBuffer.WriteString("\n")
			}
		}

		return linksBuffer.String()
	}
	return ""
}

func sortMap(routesMap map[string]Route) []Route {
	routes := make([]Route, 0, len(routesMap))
	for _, route := range routesMap {
		routes = append(routes, route)
	}
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Price < routes[j].Price
	})
	return routes
}
