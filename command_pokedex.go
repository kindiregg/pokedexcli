package main

import (
	"fmt"
)

func commandPokedex(config *Config, args ...string) error {
	if len(config.CaughtPokemon) == 0 {
		return fmt.Errorf("no pokemon caught yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
