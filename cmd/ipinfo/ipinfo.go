package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//GetLocation is a function to get the locaation of the user
func GetLocation() (location Location, err error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&location)
	if err != nil {
		return
	}

	if location.Loc == "" {
		return location, fmt.Errorf("Location cannot be determined")
	}

	latLng := strings.Split(location.Loc, ",")
	location.Lat = latLng[0]
	location.Lng = latLng[1]

	return location, nil

}
