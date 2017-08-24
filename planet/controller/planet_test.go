package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPlanetRouter(t *testing.T) {
	router := gin.Default()
	publicRoute := router.Group("/")
	Create(publicRoute)

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}

	if res.StatusCode == 404 {
		t.Errorf("expected no 404 status")
	}

	res.Body.Close()
}
