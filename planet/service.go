package planet

import (
	"errors"
	"math/rand"
	"time"

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
		return 0, errors.New("Invalid range")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Int63n(max-min), nil
}
