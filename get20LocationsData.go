package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kindiregg/pokedexcli/internal/pokecache"
)

func get20LocationsData(url string, cache pokecache.Cache) ([]byte, error) {
	if data, found := cache.Get(url); found {
		fmt.Println("Using cached data for:", url)
		return data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Store in cache for future use
	cache.Add(url, body)

	return body, nil
}
