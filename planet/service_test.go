package planet

import "testing"

func TestPickPlanet(t *testing.T) {
	planet, err := PickPlanet()
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	if planet.Climate == "" {
		t.Fatalf("expected a climate for this planet")
	}
	if planet.Population == "" {
		t.Fatalf("expected a terrain for this planet")
	}
	if planet.Name == "" {
		t.Fatalf("expected a name for this planet")
	}
	films := len(planet.FilmURLs)
	if films != planet.Films {
		t.Fatalf("expected %d count got %d", films, planet.Films)
	}
}

func TestRngInvalidRange(t *testing.T) {
	_, err := randInt64(1, 0)
	if err == nil {
		t.Fatalf("expected err count got %v", err)
	}
}
