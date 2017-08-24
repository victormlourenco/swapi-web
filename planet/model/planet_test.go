package model

import (
	"testing"
)

func TestPlanetCountRequest(t *testing.T) {
	count, err := Planet{}.Count()
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	if count < 0 {
		t.Fatalf("expected a positive count got %d", count)
	}
}

func TestPlanetRequest(t *testing.T) {
	planet := Planet{ID: 2}
	err := planet.Get()
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
}

func TestPlanetWithFilmRequest(t *testing.T) {
	planet := Planet{ID: 1}
	err := planet.Get()
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	films := len(planet.FilmURLs)
	if films != planet.Films {
		t.Fatalf("expected %d count got %d", films, planet.Films)
	}
}
