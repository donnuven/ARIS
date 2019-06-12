package ipinfo

//This for the API URL of ipinfo
const (
	apiURL = "http://ipinfo.io"
)

//Info struct
type Info struct {
	Zip     string `json:"postal"`
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Loc     string `json:"loc"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

// Location Info gives the data points of the location
type Location Info
