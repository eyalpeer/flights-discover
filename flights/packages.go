package flights

type FlightPackage struct {
	Packages []Package `json:"packages"`
}

type Package struct {
	Route []Route `json:"route"`
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
