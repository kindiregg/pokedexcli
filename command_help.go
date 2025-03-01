package main

import (
	"fmt"

	"github.com/kindiregg/pokedexcli/internal/pokecache"
)

func commandHelp(config *Config, cache *pokecache.Cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands(config, cache) {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()
	return nil
}
