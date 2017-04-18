package totoro

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Species data type
type Species struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Classification string `json:"classification"`
	EyeColor string `json:"eye_color"`
	HairColor string `json:"hair_color"`
	People []string `json:"people"`
	Films []string `json:"films"`
	URL string `json:"url"`
}

//GetSpecies gets species
func GetSpecies(query ...map[string]string) ([]Species, error){
	var species []Species

	params := url.Values{}

	if len(query) > 0 {
		for key, value := range query[0] {
			params.Set(key, value)
		}
	}

	speciesRes, err := http.Get(apiURL + "/species?" + params.Encode())

	if err != nil {
		return species, err
	}
	defer speciesRes.Body.Close()

	speciesBytes, byteErrors := ioutil.ReadAll(speciesRes.Body)

	if byteErrors != nil {
		return species, byteErrors
	}

	jsonErr := json.Unmarshal(speciesBytes, &species)

	if jsonErr != nil {
		return species, jsonErr
	}

	return species, nil
}

//GetSpeciesByID gets a species by an id
func GetSpeciesByID(id string) (Species, error) {
	speciesRes, err := http.Get(apiURL + "/species/" + id)

	if err != nil {
		return Species{}, err
	}
	defer speciesRes.Body.Close()

	speciesBytes, byteErr := ioutil.ReadAll(speciesRes.Body)

	if byteErr != nil {
		return Species{}, byteErr
	}

	var species Species

	jsonError := json.Unmarshal(speciesBytes, &species)

	if jsonError != nil {
		return Species{}, jsonError
	}

	return species, nil
}
