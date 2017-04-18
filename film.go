package totoro

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Film data type
type Film struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Director string `json:"director"`
	Producer string `json:"producer"`
	ReleaseDate string `json:"release_date"`
	RtScore string `json:"rt_score"`
	People []string `json:"people"`
	Species []string `json:"species"`
	Locations []string `json:"locations"`
	Vehicles []string `json:"vehicles"`
	URL string `json:"url"`
}

//GetFilms returns a film
func GetFilms(query ...map[string]string) ([]Film, error) {
	params := url.Values{}

	if len(query) > 0 {
		for key, value := range query[0] {
			params.Set(key, value)
		}
	}

	filmRes, err := http.Get(apiURL + "/films?" + params.Encode())

	var films []Film
	if err != nil {
		return films, err
	}
	defer filmRes.Body.Close()

	filmBytes, _ := ioutil.ReadAll(filmRes.Body)

	jsonErr := json.Unmarshal(filmBytes, &films)

	if jsonErr != nil {
		return films, jsonErr
	}
	return films, nil
}

//GetFilmByID gets a film by a specific ID
func GetFilmByID(id string) (Film, error) {
	filmRes, err := http.Get(apiURL + "/films/" + id)
	if err != nil {
		return Film{}, err
	}
	defer filmRes.Body.Close()

	filmBytes, _ := ioutil.ReadAll(filmRes.Body)

	var film Film

	jsonErr := json.Unmarshal(filmBytes, &film)

	if jsonErr != nil {
		return Film{}, err
	}
	return film, nil
}
