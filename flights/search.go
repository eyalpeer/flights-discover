package flights

type Search struct {
	Getmode         string `json:"getmode"`
	TimeFilters     []int  `json:"timefilters"`
	TypeFlight      string `json:"typeFlight"`
	Sortorder       string `json:"sortorder"`
	SortRadio       string `json:"sortradio"`
	Mode            string `json:"mode"`
	SubMode         string `json:"submode"`
	Locale          string `json:"locale"`
	Market          string `json:"market"`
	HitsLimit       int    `json:"hitslimit"`
	CalupDate       int    `json:"calupdate"`
	Cc              string `json:"cc"`
	OneForCity      string `json:"oneforcity"`
	OnePerDate      string `json:"oneperdate"`
	Currency        string `json:"currency"`
	Adults          string `json:"adults"`
	Children        string `json:"children"`
	Infants         string `json:"infants"`
	Class           string `json:"class"`
	CarryOns        int    `json:"carryons"`
	CheckedLuggages int    `json:"checkedluggages"`
	Airlines        string `json:"airlines"`
	Airports        string `json:"airports"`
	EndAirports     string `json:"endairports"`
	Stopovers       string `json:"stopovers"`
	MaxStops        int    `json:"maxstops"`
	Version         int    `json:"version"`
	Useragent       string `json:"useragent"`
	DeviceType      string `json:"devicetype"`
	ClickId         string `json:"clickid"`
	Bundle          string `json:"bundle"`
	MinPrice        int    `json:"minprice"`
	MaxPrice        int    `json:"maxprice"`
	LegList         []Leg  `json:"leglist"`
	SearchId        int    `json:"searchid"`
	UserIP          string `json:"user_ip"`
}
