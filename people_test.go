package totoro

import (
	"testing"
)

func Test_GetPeople(t *testing.T) {
	people, err := GetPeople()

	if err != nil {
		t.Error(err)
	}
	if len(people) > 10 {
		t.Log("Passed get people")
	}

	peopleLimit, limitErr := GetPeople(map[string]string{
		"query": "10",
	})

	if limitErr != nil {
		t.Error(err)
	}
	if len(peopleLimit) == 10 {
		t.Log("Passed people limit")
	}
}

func Test_GetPeopleByID(t *testing.T) {
	person, err := GetPeopleByID("ba924631-068e-4436-b6de-f3283fa848f0")

	if err != nil {
		t.Error(err)
	}

	if person.Name == "Ashitaka" {
		t.Log("Passed get person by id")
	}
}
