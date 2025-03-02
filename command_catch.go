package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *Config, args ...string) error {
	wildPokemon := args[0]

	if len(args) < 1 {
		fmt.Println("Usage: catch <name of pokemon>")
	}
	pokemonData, err := config.pokeapiClient.GetPokemon(wildPokemon)
	if err != nil {
		return fmt.Errorf("could not get pokemon '%s', not in database: %w", wildPokemon, err)

	}

	if len(config.encounteredPokemon) == 0 {
		return fmt.Errorf("you need to <explore> first to catch pokemon")
	}

	found := false

	for _, pokemon := range config.encounteredPokemon {
		if pokemon == wildPokemon {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("no %s found here, try exploring again", wildPokemon)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", wildPokemon)
	expLevel := pokemonData.BaseExperience
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := rng.Float64()
	catchProbability := 1 - (expLevel / 400)

	caught := chance <= float64(catchProbability)
	if caught {
		config.CaughtPokemon[wildPokemon] = pokemonData
		fmt.Printf("%s was caught!\n", wildPokemon)
	} else {
		fmt.Printf("%s escaped!\n", wildPokemon)
	}

	return nil

}

type PokemonResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
