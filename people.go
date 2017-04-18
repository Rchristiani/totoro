package totoro

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
)

//People data structure
type People struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age string `json:"age"`
	EyeColor string `json:"eye_color"`
	HairColor string `json:"hair_color"`
	Films []string `json:"films"`
	Species string `json:"species"`
	URL string `json:"url"`
}

//GetPeople returns a slice of people
func GetPeople(query ...map[string]string) ([]People, error){

	var people []People

	params := url.Values{}

	if len(query) > 0 {
		for key, value := range query[0] {
			params.Set(key, value)
		}
	}

	peopleRes, err := http.Get(apiURL + "/people?" + params.Encode())

	if err != nil {
		return people, err
	}
	defer peopleRes.Body.Close()

	peopleBytes, byteErr := ioutil.ReadAll(peopleRes.Body)
	if byteErr != nil {
		return people, byteErr
	}

	jsonErr := json.Unmarshal(peopleBytes, &people)

	if jsonErr != nil {
		return people, jsonErr
	}

	return people, nil
}

//GetPeopleByID gets a person by a specific id
func GetPeopleByID(id string) (People, error) {
	personRes, err := http.Get(apiURL + "/people/" + id)
	if err != nil {
		return People{}, err
	}

	personBytes, byteError := ioutil.ReadAll(personRes.Body)

	if byteError != nil {
		return People{}, byteError
	}

	var person People

	personJSONError := json.Unmarshal(personBytes, &person)

	if personJSONError != nil {
		return People{}, personJSONError
	}

	return person, nil

}
