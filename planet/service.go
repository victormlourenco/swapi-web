package planet

import (
	"errors"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	planetModel "github.com/victormlourenco/swapi-web/planet/model"
)

// PickPlanet : Picks a random planet from Swapi
func PickPlanet() (planetModel.Planet, error) {
	totalPlanets, err := planetModel.Planet{}.Count()
	if err != nil {
		return planetModel.Planet{}, err
	}
	planetID, err := randInt64(1, totalPlanets)
	if err != nil {
		return planetModel.Planet{}, err
	}
	planet := planetModel.Planet{ID: planetID}
	err = planet.Get()
	if err != nil {
		return planetModel.Planet{}, err
	}
	return planet, nil
}

func randInt64(min int64, max int64) (int64, error) {
	if (max - min) <= 0 {
		err := errors.New("Invalid range")
		log.WithFields(log.Fields{
			"Min": min,
			"Max": max,
		}).Error(err)
		return 0, err
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Int63n(max-min), nil
}
