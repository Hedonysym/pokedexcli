package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokecache "github.com/Hedonysym/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	cache := pokecache.NewCache(5 * 1000)
	config := &Config{
		prevUrl: "",
		nextUrl: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		success := scanner.Scan()
		if !success {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
			}
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		if cmd, exists := commandRegistry[input[0]]; !exists {
			fmt.Print("Unknown command\n")
		} else {
			if err := cmd.callback(config, cache); err != nil {
				fmt.Fprintln(os.Stderr, "Error executing command:", err)
			}
		}
	}
}
