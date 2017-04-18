package totoro

import (
	"testing"
)

func Test_GetFilms(t *testing.T) {
	films, err := GetFilms()
	if err != nil {
		t.Error(err)
	}
	if len(films) > 10 {
		t.Log("Passed get films")
	}

	filmLimit, limitErr := GetFilms(map[string]string{
		"limit": "10",
	})
	if limitErr != nil {
		t.Error(limitErr)
	}

	if len(filmLimit) == 10 {
		t.Log("Passed get films limit 10")
	}

}

func Test_GetFilmByID(t *testing.T) {
	film, err := GetFilmByID("2baf70d1-42bb-4437-b551-e5fed5a87abe")

	if err != nil {
		t.Error(err)
	}
	if film.Title == "Castle in the Sky" {
		t.Log("Passed get film by id")
	}
}
