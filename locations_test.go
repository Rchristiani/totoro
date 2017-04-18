package totoro

import (
	"testing"
)

func Test_GetLocations(t *testing.T) {
	locations, err := GetLocations()
	if err != nil {
		t.Error(err)
	}

	if len(locations) > 10 {
		t.Log("Passed geting locations")
	} else {
		t.Error("Failed to get locations")
	}

	//Skipping for now, limit doesn't seem to work on this endpoint

	locationLimit, limitErr := GetLocations(map[string]string{
		"limit": "5",
	})
	t.Skip("Skipping the get location by limit, currently not working from API")
	if limitErr != nil {
		t.Error(limitErr)
	}

	if len(locationLimit) == 5 {
		t.Log("Passed get locations with limit")
	} else {
		t.Error("Failed to get locations with limit")
	}

}

func Test_GetLocationByID(t *testing.T) {
	location, err := GetLocationByID("11014596-71b0-4b3e-b8c0-1c4b15f28b9a")

	if err != nil {
		t.Error(err)
	}

	if location.Name == "Irontown" {
		t.Log("Passed get location by id")
	} else {
		t.Error("Failed to get location by id")
	}
}
