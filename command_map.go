package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != "" {
		url = config.Next
	}
	err := fetchAndDisplayLocations(url, config)
	if err != nil {
		return err
	}
	return nil
}

func commandMapBack(config *Config) error {
	if config.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	err := fetchAndDisplayLocations(config.Previous, config)
	if err != nil {
		return err
	}
	return nil
}

// Helper function to fetch and display locations
func fetchAndDisplayLocations(url string, config *Config) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not get locations: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response")
	}

	var locationResp LocationAreaResponse
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return fmt.Errorf("could not parse response")
	}

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	if locationResp.Next != nil {
		config.Next = *locationResp.Next
	} else {
		config.Next = ""
	}

	if locationResp.Previous != nil {
		config.Previous = *locationResp.Previous
	} else {
		config.Previous = ""
	}

	return nil
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}
