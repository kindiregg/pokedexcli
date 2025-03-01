package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(commands map[string]cliCommand) {
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

		if command, ok := commands[commandName]; ok {
			if err := command.callback(); err != nil {
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *Config
}

func getCommands(config *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return commandExit(config) },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(config) },
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    func() error { return commandMap(config) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    func() error { return commandMapBack(config) },
		},
	}
}

type Config struct {
	Next     string
	Previous string
}
