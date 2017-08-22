package planet

import (
	"math/rand"
	"time"

	planetModel "github.com/swapi-web/planet/model"
)

// PickPlanet : Picks a random planet from Swapi
func PickPlanet() (planetModel.Planet, error) {
	totalPlanets, err := planetModel.Planet{}.Count()
	if err != nil {
		return planetModel.Planet{}, err
	}
	planet := planetModel.Planet{ID: randInt64(1, totalPlanets)}
	err = planet.Get()
	if err != nil {
		return planetModel.Planet{}, err
	}
	return planet, nil
}

func randInt64(min int64, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return min + rand.Int63n(max-min)
}
