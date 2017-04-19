package totoro

import(
	"testing"
)

func Test_GetVehicles(t *testing.T) {
	vehicles, err := GetVehicles()

	if err != nil {
		t.Error(err)
	}

	if len(vehicles) > 0 {
		t.Log("Passed get vehicles")
	}

}

func Test_GetVehiclesByID(t *testing.T) {
	vehcile, err := GetVehiclesByID("923d70c9-8f15-4972-ad53-0128b261d628")

	if err != nil {
		t.Error(err)
	}

	if vehcile.Name == "Sosuke's Boat" {
		t.Log("Passed get vehicle by ID")
	} else {
		t.Error(err)
	}
}
