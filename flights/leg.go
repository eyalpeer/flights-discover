package flights

type Leg struct {
	TypeDirection   string  `json:"typedirection"`
	FromDisplayName string  `json:"fromdisplayname"`
	ToDisplayName   string  `json:"todisplayname"`
	FromLocLat      float64 `json:"fromloclat"`
	FromLocLng      float64 `json:"fromloclng"`
	FromName        string  `json:"fromname"`
	ToLocLat        float64 `json:"toloclat"`
	ToLocLng        float64 `json:"toloclng"`
	ToName          string  `json:"toname"`
	FromLocRad      string  `json:"fromlocrad"`
	FromLocSid      string  `json:"fromlocsid"`
	ToLocRad        string  `json:"tolocrad"`
	ToLocRadURL     string  `json:"tolocrad_URL"`
	ToLocSid        string  `json:"tolocsid"`
	Dffy            string  `json:"dffy"`
	Dffm            string  `json:"dffm"`
	Dffd            string  `json:"dffd"`
	Dfty            string  `json:"dfty"`
	Dftm            string  `json:"dftm"`
	Dftd            string  `json:"dftd"`
	Dtfy            string  `json:"dtfy"`
	Dtfm            string  `json:"dtfm"`
	Dtfd            string  `json:"dtfd"`
	Dtty            string  `json:"dtty"`
	Dttm            string  `json:"dttm"`
	Dttd            string  `json:"dttd"`
	Somin           int     `json:"somin"`
	Somax           int     `json:"somax"`
}
