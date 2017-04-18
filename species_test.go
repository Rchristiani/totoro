package totoro

import (
	"testing"
)

func Test_GetSpecies(t *testing.T) {
	species, err := GetSpecies()

	if err != nil {
		t.Error(err)
	}

	if len(species) > 0 {
		t.Log("Passed get species")
	} else {
		t.Error("Failed to get species")
	}

	limitSpecies, limitErr := GetSpecies(map[string]string{
		"limit": "1",
	})
	t.Skip("Skipping this test since it seems to not work from the API")
	if limitErr != nil {
		t.Error("Failed to get species with a specific limit")
	}

	if len(limitSpecies) == 2 {
		t.Log("Passed getting species with a limit")
	} else {
		t.Error("Failed to get species with a limit")
	}

}

func Test_GetSpeciesByID(t *testing.T) {
	species, err := GetSpeciesByID("b5a92d0e-5fb4-43d4-ba60-c012135958e4")

	if err != nil {
		t.Error("Failed get species by id")
	}

	if species.Name == "Spirit" {
		t.Log("Passed get species by id")
	} else {
		t.Error("Failed to get species by id, name wrong")
	}
}
