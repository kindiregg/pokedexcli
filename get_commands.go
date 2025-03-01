package main

import (
	"github.com/kindiregg/pokedexcli/internal/pokecache"
)

func getCommands(config *Config, cache *pokecache.Cache) map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(params ...string) error { return commandExit() },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(params ...string) error { return commandHelp(config, cache) },
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    func(params ...string) error { return commandMap(config, cache) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    func(params ...string) error { return commandMapBack(config, cache) },
		},
		"explore": {
			name:        "explore",
			description: "explores area by name or id",
			callback:    func(params ...string) error { return commandExplore(config, cache, params...) },
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(params ...string) error
}

type Config struct {
	Next     string
	Previous string
}
