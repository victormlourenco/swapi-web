package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/victormlourenco/swapi-web/config"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

// Planet : Planet object
type Planet struct {
	ID         int64
	Name       string   `json:"name"`
	Climate    string   `json:"climate"`
	Terrain    string   `json:"terrain"`
	Population string   `json:"population"`
	FilmURLs   []string `json:"films"`
	Films      int
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

	if req.StatusCode != http.StatusOK {
		err := fmt.Errorf("The API server returned status code %d", req.StatusCode)
		log.WithFields(log.Fields{
			"RequestURL": req.Request.URL,
		}).Error(err)
		return 0, err
	}

	planets := Planets{}

	err = json.NewDecoder(req.Body).Decode(&planets)
	if err != nil {
		log.WithFields(log.Fields{
			"RequestURL": req.Request.URL,
		}).Error(err)
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

	if req.StatusCode != http.StatusOK {
		err := fmt.Errorf("The API server returned status code %d", req.StatusCode)
		log.WithFields(log.Fields{
			"RequestURL": req.Request.URL,
			"PlanetID":   p.ID,
		}).Error(err)
		return err
	}

	err = json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		log.WithFields(log.Fields{
			"RequestURL": req.Request.URL,
			"PlanetID":   p.ID,
		}).Error(err)
		return err
	}

	p.Films = len(p.FilmURLs)

	return nil
}
