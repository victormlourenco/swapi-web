package model

import (
	"encoding/json"
	"net/http"
)

var client = &http.Client{}

// Film : Film object
type Film struct {
	Title string `json:"title"`
	URL   string
}

// Get : Get a Film by it's URL
func (f *Film) Get() error {
	req, err := client.Get(f.URL)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(f)
	if err != nil {
		return err
	}

	return nil
}
