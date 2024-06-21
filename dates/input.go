package dates

type Input struct {
	DepartureStartYear  string `json:"dffy"`
	DepartureStartMonth string `json:"dffm"`
	DepartureStartDay   string `json:"dffd"`
	DepartureEndYear    string `json:"dfty"`
	DepartureEndMonth   string `json:"dftm"`
	DepartureEndDay     string `json:"dftd"`
	ArriveStartYear     string `json:"dtfy"`
	ArriveStartMonth    string `json:"dtfm"`
	ArriveStartDay      string `json:"dtfd"`
	ArriveEndYear       string `json:"dtty"`
	ArriveEndMonth      string `json:"dttm"`
	ArriveEndDay        string `json:"dttd"`
}
