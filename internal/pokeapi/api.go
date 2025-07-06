package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationPage struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetMapPage(url string) (LocationPage, error) {
	var fail LocationPage
	res, err := http.Get(url)
	if err != nil {
		return fail, fmt.Errorf("failed to fetch locations: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fail, fmt.Errorf("failed to fetch locations: %s", res.Status)
	}
	var page LocationPage
	if err := json.NewDecoder(res.Body).Decode(&page); err != nil {
		return fail, fmt.Errorf("failed to decode locations: %w", err)
	}
	if page.Count == 0 {
		return fail, fmt.Errorf("no locations found")
	}
	return page, nil
}
