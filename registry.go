package main

import (
	"fmt"
	"math/rand"
	"os"

	pokeapi "github.com/Hedonysym/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args []string, config *Config, client *pokeapi.Client) error
}

type Config struct {
	prevUrl string
	nextUrl string
}

var commandRegistry map[string]cliCommand

func commandExit(args []string, c *Config, client *pokeapi.Client) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(args []string, c *Config, client *pokeapi.Client) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(args []string, c *Config, client *pokeapi.Client) error {
	page, err := pokeapi.GetMapPage(c.nextUrl, client)
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

func commandMapb(args []string, c *Config, client *pokeapi.Client) error {
	if c.prevUrl == "" {
		return fmt.Errorf("no previous page available")
	}
	if c.prevUrl == c.nextUrl {
		return fmt.Errorf("already at the first page")
	}
	page, err := pokeapi.GetMapPage(c.prevUrl, client)
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

func commandExplore(args []string, c *Config, client *pokeapi.Client) error {
	if len(args) == 0 {
		return fmt.Errorf("no area name provided")
	}
	areaName := args[0]
	loc, err := pokeapi.GetLocationFull(areaName, client)
	if err != nil {
		return fmt.Errorf("failed to fetch location: %w", err)
	}
	fmt.Printf("Exploring %s\n", loc.Name)
	if len(loc.PokemonEncounters) == 0 {
		fmt.Printf("No Pokemon found in this area.\n")
	} else {
		fmt.Printf("Pokemon found in this area:\n")
		for _, encounter := range loc.PokemonEncounters {
			fmt.Printf("- %s\n", encounter.Pokemon.Name)
		}
	}
	return nil
}

func commandCatch(args []string, c *Config, client *pokeapi.Client) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name provided")
	}
	pkmnName := args[0]
	pkmn, err := pokeapi.GetPokemon(pkmnName, client)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon: %w", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pkmn.Name)
	if pkmn.BaseExperience > 400 {
		chance := rand.Intn(100)
		if chance < 70 {
			return fmt.Errorf("the %s escaped!", pkmn.Name)
		}
	} else if pkmn.BaseExperience > 200 {
		chance := rand.Intn(100)
		if chance < 50 {
			return fmt.Errorf("the %s escaped!", pkmn.Name)
		} else {
			chance := rand.Intn(100)
			if chance < 20 {
				return fmt.Errorf("the %s escaped!", pkmn.Name)
			}
		}
	}
	fmt.Printf("You caught a %s!\n", pkmn.Name)
	(*client.Pkmn)[pkmn.Name] = pkmn
	return nil
}

func commandInspect(args []string, c *Config, client *pokeapi.Client) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name provided")
	}
	pkmnName := args[0]
	pkmn, ok := (*client.Pkmn)[pkmnName]
	if !ok {
		return fmt.Errorf("you don't have a %s", pkmnName)
	}
	fmt.Printf("Name: %s\n", pkmn.Name)
	fmt.Printf("Height: %d\n", pkmn.Height)
	fmt.Printf("Weight: %d\n", pkmn.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pkmn.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pkmn.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(args []string, c *Config, client *pokeapi.Client) error {
	if len(*client.Pkmn) == 0 {
		fmt.Print("You haven't caught any Pokemon yet.\n")
		return nil
	}
	fmt.Print("Caught Pokemon:\n")
	for name := range *client.Pkmn {
		fmt.Printf("- %s\n", name)
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
		"explore": {
			name:        "explore <area-name>",
			description: "Explore a location to find Pokemon, use the name as shown in the map command.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch a Pokemon by name, as shown in the explore command.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect a caught Pokemon by name.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon.",
			callback:    commandPokedex,
		},
	}
}
