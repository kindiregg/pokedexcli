package main

import (
	"github.com/kindiregg/pokedexcli/internal/pokecache"
)

func getCommands(config *Config, cache *pokecache.Cache) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return commandExit() },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(config, cache) },
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    func() error { return commandMap(config, cache) },
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    func() error { return commandMapBack(config, cache) },
		},
	}
}
