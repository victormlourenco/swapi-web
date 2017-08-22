package model

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/swapi-web/config"

	filmModel "github.com/swapi-web/film/model"
)

var client = &http.Client{}

// Planet : Planet object
type Planet struct {
	ID         int64
	Name       string   `json:"name"`
	Climate    string   `json:"climate"`
	Terrain    string   `json:"terrain"`
	Population string   `json:"population"`
	FilmURLs   []string `json:"films"`
	Films      []filmModel.Film
}

// Planets : Group of planets
type Planets struct {
	Count int64 `json:"count"`
}

// Count : Get the total of planets from Swapi
func (p Planet) Count() (int64, error) {
	req, err := client.Get(fmt.Sprintf("%s/planets/", config.GetString("RemoteUrl")))
	if err != nil {
		return 0, err
	}
	defer req.Body.Close()

	planets := Planets{}

	err = json.NewDecoder(req.Body).Decode(&planets)
	if err != nil {
		return 0, err
	}

	return planets.Count, nil
}

// Get : Get a Planet by it's ID
func (p *Planet) Get() error {
	req, err := client.Get(fmt.Sprintf("%s/planets/%d", config.GetString("RemoteUrl"), p.ID))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		return err
	}

	for k := range p.FilmURLs {
		film := filmModel.Film{URL: p.FilmURLs[k]}
		err := film.Get()
		if err != nil {
			return err
		}
		p.Films = append(p.Films, film)
	}

	return nil
}
