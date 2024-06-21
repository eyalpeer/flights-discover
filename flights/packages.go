package flights

import (
	"bytes"
	"fmt"
)

type FlightPackage struct {
	Packages []Package `json:"packages"`
}

type Package struct {
	Route []Route `json:"route"`
}

type Route struct {
	Route              []Segment   `json:"route"`
	Conversion         interface{} `json:"conversion"`
	DeepLink           string      `json:"deep_link"`
	Baggage            Baggage     `json:"baggage"`
	Price              float64     `json:"price"`
	FlyDuration        string      `json:"fly_duration"`
	FlyDurationSecs    float64     `json:"fly_duration_secs"`
	FlyFrom            string      `json:"flyFrom"`
	FlyTo              string      `json:"flyTo"`
	DTime              int64       `json:"dTime"`
	ATime              int64       `json:"aTime"`
	ReturnDuration     string      `json:"return_duration"`
	ReturnDurationSecs float64     `json:"return_duration_secs"`
	NightsInDest       interface{} `json:"nightsInDest"`
	CityCodeFrom       string      `json:"cityCodeFrom"`
	CityCodeTo         string      `json:"cityCodeTo"`
	OperatingCarrier   string      `json:"operating_carrier"`
	Sc                 string      `json:"sc"`
	ArrivalPos         int         `json:"arrivalpos"`
	TripSpec           []string    `json:"tripspec"`
	Airports           string      `json:"airports"`
	EndAirports        string      `json:"endairports"`
	Airlines           string      `json:"airlines"`
	MaxStops           int         `json:"maxstops"`
	Co2                int         `json:"co2"`
	FlyDurationMin     float64     `json:"fly_duration_min"`
	SegDurMax          float64     `json:"segdurmax"`
	DrTime             int64       `json:"drTime"`
	ArTime             int64       `json:"arTime"`
	ReturnDurationMin  float64     `json:"return_duration_min"`
}

type Segment struct {
	FlyFrom          string `json:"flyFrom"`
	FlyTo            string `json:"flyTo"`
	DTime            int    `json:"dTime"`
	ATime            int    `json:"aTime"`
	Cabin            string `json:"cabin"`
	Sc               string `json:"sc"`
	OperatingCarrier string `json:"operating_carrier"`
	Airline          string `json:"airline"`
	FlightNo         int    `json:"flight_no"`
	Return           int    `json:"return"`
	AirlineName      string `json:"airlinename"`
	LocalDep         string `json:"local_dep"`
}

type Baggage struct {
	Unit   string  `json:"unit"`
	Amount int     `json:"amount"`
	Bin    float64 `json:"bin"`
}

func (fp FlightPackage) String() string {
	var buffer bytes.Buffer
	for _, p := range fp.Packages {
		for _, r := range p.Route {
			for _, s := range r.Route {
				buffer.WriteString(fmt.Sprintf("%s -> %s\n", s.FlyFrom, s.FlyTo))
			}
		}
	}
	return buffer.String()
}
