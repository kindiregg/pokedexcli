package main

import (
	"encoding/json"
	"fmt"

	"github.com/kindiregg/pokedexcli/internal/pokecache"
)

func commandExplore(config *Config, cache *pokecache.Cache, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("no location provided, please provide a location")
	}

	locationName := params[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationName)

	data, err := getAPIData(url, cache)
	if err != nil {
		return fmt.Errorf("could not get location data: %w", err)
	}

	var locationDetail LocationAreaDetail

	if err := json.Unmarshal(data, &locationDetail); err != nil {
		return err
	}

	fmt.Printf("Exploring %s... \n", locationDetail.Name)

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

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
