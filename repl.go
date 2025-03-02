package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kindiregg/pokedexcli/internal/pokeapi"
)

type Config struct {
	Next          string
	Previous      string
	caughtPokemon map[string]pokeapi.Pokemon
	pokeapiClient pokeapi.Client
}

func startRepl(config *Config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			continue
		}
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		args := cleanedInput[1:]

		if command, ok := commands[commandName]; ok {
			if err := command.callback(config, args...); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}
