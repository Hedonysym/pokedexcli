package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Hedonysym/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	client := pokeapi.NewClient(5 * 1000)
	config := &Config{
		prevUrl: "",
		nextUrl: "",
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
			if err := cmd.callback(input[1:], config, client); err != nil {
				fmt.Fprintln(os.Stderr, "Error executing command:", err)
			}
		}
	}
}
