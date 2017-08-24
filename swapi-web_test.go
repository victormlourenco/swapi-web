package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/victormlourenco/swapi-web/config"
)

func TestHttpHandler(t *testing.T) {
	go main()

	client := &http.Client{}

	req, err := client.Get(fmt.Sprintf("http://%s/", config.GetString("ServerAddress")))
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		t.Fatalf("handler returned wrong status code: got %d want %d", req.StatusCode, http.StatusOK)
	}

	req, err = client.Get(fmt.Sprintf("http://%s/assets/images/favicon.png", config.GetString("ServerAddress")))
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		t.Fatalf("handler returned wrong status code: got %d want %d", req.StatusCode, http.StatusOK)
	}

	req, err = client.Get(fmt.Sprintf("http://%s/assets/css/styles.css", config.GetString("ServerAddress")))
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		t.Fatalf("handler returned wrong status code: got %d want %d", req.StatusCode, http.StatusOK)
	}
}
