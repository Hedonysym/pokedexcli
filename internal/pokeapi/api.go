package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func GetMapPage(url string, client *Client) (LocationPage, error) {
	var fail LocationPage
	if url == "" {
		url = baseURL + "/location-area/?offset=0&limit=20"
	}
	val, ok := client.cache.Get(url)
	if ok {
		var page LocationPage
		if err := json.Unmarshal(val, &page); err != nil {
			return fail, fmt.Errorf("failed to decode locations: %w", err)
		}
		if page.Count == 0 {
			return fail, fmt.Errorf("no locations found")
		}
		return page, nil
	}
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

func GetLocationFull(url string, client *Client) (LocationFull, error) {
	var fail LocationFull
	if url == "" {
		return fail, fmt.Errorf("no name provided")
	}
	fullurl := baseURL + "/location-area/" + url + "/"
	val, ok := client.cache.Get(fullurl)
	if ok {
		var loc LocationFull
		if err := json.Unmarshal(val, &loc); err != nil {
			return fail, fmt.Errorf("failed to decode location: %w", err)
		}
		if loc.ID == 0 {
			return fail, fmt.Errorf("location not found")
		}
		return loc, nil
	}
	res, err := http.Get(fullurl)
	if err != nil {
		return fail, fmt.Errorf("failed to fetch location: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fail, fmt.Errorf("failed to fetch location: %s", res.Status)
	}
	var loc LocationFull
	if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
		return fail, fmt.Errorf("failed to decode location: %w", err)
	}
	if loc.ID == 0 {
		return fail, fmt.Errorf("location not found")
	}
	return loc, nil

}

func GetPokemon(url string, client *Client) (Pokemon, error) {
	var fail Pokemon
	if url == "" {
		return fail, fmt.Errorf("no name provided")
	}
	fullurl := baseURL + "/pokemon/" + url + "/"
	val, ok := (*client.Pkmn)[fullurl]
	if ok {
		return val, nil
	}
	res, err := http.Get(fullurl)
	if err != nil {
		return fail, fmt.Errorf("failed to fetch pokemon: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fail, fmt.Errorf("failed to fetch pokemon: %s", res.Status)
	}
	var pkmn Pokemon
	if err := json.NewDecoder(res.Body).Decode(&pkmn); err != nil {
		return fail, fmt.Errorf("failed to decode pokemon: %w", err)
	}
	if pkmn.ID == 0 {
		return fail, fmt.Errorf("pokemon not found")
	}
	return pkmn, nil
}
