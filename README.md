# Totoro

[![CircleCI](https://img.shields.io/circleci/project/Rchristiani/bark.svg?style=flat-square)](https://circleci.com/gh/Rchristiani/bark)

A Go wrapper for the Studio Ghibli API (https://ghibliapi.herokuapp.com)[ https://ghibliapi.herokuapp.com]

### Usage

```go
package main

import (
	"fmt"
	"github.com/rchristiani/totoro" //import package
)

func main() {
	people, err := totoro.GetPeople()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people)
}
```

The the `totoro` package contains methods for all endpoints.

## Films

 The documentation can be found [here](https://ghibliapi.herokuapp.com/#tag/Films).

#### Film struct

```go
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
```

### `GetFilms(query ...map[string]string) ([]Film, error)`

Returns a slice of `Film`, you can pass a query that can be used to set the `limit` param.

```go
films, err := totoro.GetFilms(map[string]string{
	"limit":"10",
})
```

### `GetFilmByID(id string) (Film, error)`

Returns a single `Film` by `id`.

## People

The documentation can be found [here](https://ghibliapi.herokuapp.com/#tag/People).


#### People struct

```go
ID string `json:"id"`
Name string `json:"name"`
Gender string `json:"gender"`
Age string `json:"age"`
EyeColor string `json:"eye_color"`
HairColor string `json:"hair_color"`
Films []string `json:"films"`
Species string `json:"species"`
URL string `json:"url"`
```

### `GetPeople(query ...map[string]string) ([]People, error)`

Returns a slice of `People`, you can pass a query that can be used to set the `limit` param.

```go
films, err := totoro.GetPeople(map[string]string{
	"limit":"10",
})
```

### `GetPeopleByID(id string) (Film, error)`

Returns a single `People` by `id`.

## Locations

The documentation can be found [here](https://ghibliapi.herokuapp.com/#tag/Locations).

#### Location struct

```go
ID string `json:"id"`
Name string `json:"name"`
Climate string `json:"climate"`
Terrain string `json:"terrain"`
SurfaceWater string `json:"surface_water"`
Residents []string `json:"residents"`
Films []string `json:"films"`
URL []string `json:"url"`
```

### `GetLocations(query ...map[string]string) ([]Location, error)`

Returns a slice of `Location`, you can pass a query that can be used to set the `limit` param.

```go
films, err := totoro.GetLocations(map[string]string{
	"limit":"10",
})
```

### `GetLocationByID(id string) (Location, error)`

Returns a single `Location` by `id`.

## Species

The documentation can be found [here](https://ghibliapi.herokuapp.com/#tag/Species).

#### Species struct

```go
ID string `json:"id"`
Name string `json:"name"`
Classification string `json:"classification"`
EyeColor string `json:"eye_color"`
HairColor string `json:"hair_color"`
People []string `json:"people"`
Films []string `json:"films"`
URL string `json:"url"`
```

### `GetSpecies(query ...map[string]string) ([]Species, error)`

Returns a slice of `Species`, you can pass a query that can be used to set the `limit` param.

```go
films, err := totoro.GetSpecies(map[string]string{
	"limit":"10",
})
```

### `GetSpeciesByID(id string) (Species, error)`

Returns a single `Species` by `id`.


## Vehicles

The documentation can be found [here](https://ghibliapi.herokuapp.com/#tag/Species).

#### Vehicles struct

```go
ID string `json:"id"`
Name string `json:"name"`
Description string `json:"description"`
VehicleClass string `json:"vehicle_class"`
Length string `json:"length"`
Pilot string `json:"pilot"`
Films string `json:"films"`
URL string `json:"url"`
```

### `GetVehicles(query ...map[string]string) ([]Vehicles, error)`

Returns a slice of `Vehicles`, you can pass a query that can be used to set the `limit` param.

```go
films, err := totoro.GetVehicles(map[string]string{
	"limit":"10",
})
```

### `GetVehiclesByID(id string) (Vehicles, error)`

Returns a single `Vehicles` by `id`.
