package main

import (
	"github.com/kindiregg/pokedexcli/internal/pokecache"
	"time"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	println(cache)
	config := &Config{}
	commands := getCommands(config, cache)
	startRepl(commands)
}
