package main

import (
	"fmt"
	"os"

	pokecache "github.com/Hedonysym/pokedexcli/internal/pokecache"

	pokeapi "github.com/Hedonysym/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, cache *pokecache.Cache) error
}

type Config struct {
	prevUrl string
	nextUrl string
}

var commandRegistry map[string]cliCommand

func commandExit(c *Config, cache *pokecache.Cache) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, cache *pokecache.Cache) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(c *Config, cache *pokecache.Cache) error {
	page, err := pokeapi.GetMapPage(c.nextUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch map page: %w", err)
	}
	c.prevUrl = page.Previous
	c.nextUrl = page.Next
	for _, loc := range page.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

func commandMapb(c *Config, cache *pokecache.Cache) error {
	if c.prevUrl == "" {
		return fmt.Errorf("no previous page available")
	}
	if c.prevUrl == c.nextUrl {
		return fmt.Errorf("already at the first page")
	}
	page, err := pokeapi.GetMapPage(c.prevUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch map page: %w", err)
	}
	c.prevUrl = page.Previous
	c.nextUrl = page.Next
	for _, loc := range page.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

func init() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp},
		"map": {
			name:        "map",
			description: "Displays the next page of locations in the Pokemon world. 20 per page.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the last page of locations in the Pokemon world. 20 per page.",
			callback:    commandMapb,
		},
	}
}
