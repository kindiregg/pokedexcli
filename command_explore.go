package main

import (
	"fmt"
	"strings"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location provided, please provide a location")
	}

	inputName := args[0]

	location, err := config.pokeapiClient.GetLocation(inputName)
	if err != nil {
		return fmt.Errorf("unable to get location %s: %w", inputName, err)
	}

	fmt.Printf("Exploring %s... \n", location.Name)

	var currentPokemon []string

	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
		currentPokemon = append(currentPokemon, strings.TrimSpace(strings.ToLower(encounter.Pokemon.Name)))
	}

	config.encounteredPokemon = currentPokemon

	return nil

}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type LocationAreaDetail struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
