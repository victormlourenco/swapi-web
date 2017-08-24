package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpHandler(t *testing.T) {
	ts := httptest.NewServer(getMainEngine())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("handler returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}
	res.Body.Close()

	res, err = http.Get(fmt.Sprintf("%s/assets/images/favicon.png", ts.URL))
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("handler returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}
	res.Body.Close()

	res, err = http.Get(fmt.Sprintf("%s/assets/css/styles.css", ts.URL))
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("handler returned wrong status code: got %d want %d", res.StatusCode, http.StatusOK)
	}
	res.Body.Close()
}
