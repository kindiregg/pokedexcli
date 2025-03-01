package main

import (
	"time"

	"github.com/kindiregg/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	config := &Config{}
	commands := getCommands(config, cache)
	startRepl(commands)
}
