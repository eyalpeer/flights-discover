package flights

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"main.go/dates"
	"net/http"
)

const URL = "https://worka.panflights.com/skypickersearchsingle"

var (
	httpClient = &http.Client{}
)

func buildRequest(dates dates.Range) (*Search, error) {
	datesEnriched, err := dates.DateRangeToInput(1)
	if err != nil {
		return nil, err
	}
	return &Search{
		Getmode: "searchflights",
		TimeFilters: []int{
			0,
			24,
			0,
			24,
			0,
			24,
			0,
			24,
		},
		TypeFlight:      "round",
		Sortorder:       "price",
		SortRadio:       "price",
		Mode:            "search",
		SubMode:         "",
		Locale:          "en",
		Market:          "il", // consider changing
		HitsLimit:       500,
		CalupDate:       0,
		Cc:              "0",
		OneForCity:      "0",
		OnePerDate:      "0",
		Currency:        "USD",
		Adults:          "2",
		Children:        "0",
		Infants:         "0",
		Class:           "Y",
		CarryOns:        0,
		CheckedLuggages: 0,
		Airlines:        "",
		Airports:        "",
		EndAirports:     "",
		Stopovers:       "",
		MaxStops:        0,
		Version:         0,
		Useragent:       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		DeviceType:      "PC",
		ClickId:         "", // consider change
		Bundle:          "{}",
		MinPrice:        0,
		MaxPrice:        999999999999,
		LegList: []Leg{
			{
				TypeDirection:   "unset",
				FromDisplayName: "Israel",
				ToDisplayName:   "Europe",
				FromLocSid:      "IL",
				ToLocRad:        "0",
				FromLocRad:      "0",
				ToLocSid:        "0",
				ToLocRadURL:     "0",
				Dffy:            datesEnriched.DepartureStartYear,
				Dffm:            datesEnriched.DepartureStartMonth,
				Dffd:            datesEnriched.DepartureStartDay,
				Dfty:            datesEnriched.DepartureEndYear,
				Dftm:            datesEnriched.DepartureEndMonth,
				Dftd:            datesEnriched.DepartureEndDay,
				Dtfy:            datesEnriched.DepartureStartYear,
				Dtfm:            datesEnriched.DepartureStartMonth,
				Dtfd:            datesEnriched.DepartureStartDay,
				Dtty:            datesEnriched.DepartureEndYear,
				Dttm:            datesEnriched.DepartureEndMonth,
				Dttd:            datesEnriched.DepartureEndDay,
				Somin:           0,
				Somax:           96,
			},
		},
	}, nil
}

func SendRequest(dates dates.Range) (*FlightPackage, error) {
	request, err := buildRequest(dates)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequest("POST", URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	response, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		errormessage, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
		}
		return nil, fmt.Errorf("unexpected status code: %d, error message: %s", response.StatusCode, errormessage)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("Failed to close response body: %v", err)
		}
	}(response.Body)
	var flights FlightPackage
	err = json.NewDecoder(response.Body).Decode(&flights)
	return &flights, err
}
