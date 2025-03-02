package main

import (
	"fmt"

	"github.com/kindiregg/pokedexcli/internal/pokeapi"
)

func getTestPokemon(c pokeapi.Client) {
	testPokemon := []string{"pikachu", "charizard", "magikarp", "mewtwo", "bulbasaur"}
	for _, name := range testPokemon {
		// Fetch Pokemon and print its base experience
		pokemon, err := c.GetPokemon(name)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", name, err)
			continue
		}
		fmt.Printf("%s base experience: %d\n", name, pokemon.BaseExperience)
	}
}
