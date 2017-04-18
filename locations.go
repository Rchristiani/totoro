package totoro

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
)

//Location data type
type Location struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
	SurfaceWater string `json:"surface_water"`
	Residents []string `json:"residents"`
	Films []string `json:"films"`
	URL []string `json:"url"`
}

//GetLocations gets all locations
func GetLocations(query ...map[string]string) ([]Location, error) {

	var locations []Location

	params := url.Values{}

	if len(query) > 0 {
		for key, value := range query[0] {
			params.Set(key, value)
		}
	}

	locationRes, err := http.Get(apiURL + "/locations?" + params.Encode())

	if err != nil {
		return locations, err
	}
	defer locationRes.Body.Close()

	locationBytes, bytesError := ioutil.ReadAll(locationRes.Body)

	if bytesError != nil {
		return locations, err
	}

	jsonError := json.Unmarshal(locationBytes, &locations)

	if jsonError != nil {
		return locations, jsonError
	}

	return locations, nil
}

//GetLocationByID gets a location by an id
func GetLocationByID(id string) (Location, error) {
	locationRes, err := http.Get(apiURL + "/locations/" + id)

	if err != nil {
		return Location{}, err
	}

	defer locationRes.Body.Close()

	locationBytes, bytesError := ioutil.ReadAll(locationRes.Body)

	if bytesError != nil {
		return Location{}, bytesError
	}

	var location Location

	jsonError := json.Unmarshal(locationBytes, &location)

	if jsonError != nil {
		return Location{}, jsonError
	}

	return location, err

}
