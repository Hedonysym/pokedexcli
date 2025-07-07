package pokeapi

import (
	"time"

	pokecache "github.com/Hedonysym/pokedexcli/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
	Pkmn  *map[string]Pokemon
}

func NewClient(cacheInterval int) *Client {
	return &Client{
		cache: pokecache.NewCache(time.Duration(cacheInterval)),
		Pkmn:  &map[string]Pokemon{},
	}
}
