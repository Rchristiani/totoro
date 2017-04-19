package totoro

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
)

//Vehicle data type
type Vehicle struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	VehicleClass string `json:"vehicle_class"`
	Length string `json:"length"`
	Pilot string `json:"pilot"`
	Films string `json:"films"`
	URL string `json:"url"`
}

//GetVehicles gets all the vehicles
func GetVehicles() ([]Vehicle, error) {

		var vehicles []Vehicle

		vehicleRes, err := http.Get(apiURL + "/vehicles")

		if err != nil {
			return vehicles, err
		}

		defer vehicleRes.Body.Close()

		vehiclesBytes, byteErrors := ioutil.ReadAll(vehicleRes.Body)

		if byteErrors != nil {
			return vehicles, byteErrors
		}

		jsonError := json.Unmarshal(vehiclesBytes, &vehicles)

		if jsonError != nil {
			return vehicles, jsonError
		}

		return vehicles, nil

}

//GetVehiclesByID gets a vehicle by id
func GetVehiclesByID(id string) (Vehicle, error){
	vehicleRes, err := http.Get(apiURL + "/vehicles/" + id)

	if err != nil {
		return Vehicle{}, err
	}

	defer vehicleRes.Body.Close()

	vehicleBytes, byteError := ioutil.ReadAll(vehicleRes.Body)

	if byteError != nil {
		return Vehicle{}, byteError
	}

	var vehicle Vehicle

	jsonError := json.Unmarshal(vehicleBytes, &vehicle)

	if jsonError != nil {
		return Vehicle{}, jsonError
	}

	return vehicle, nil
}
