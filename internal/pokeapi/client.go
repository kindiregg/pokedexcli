package pokeapi

import (
	"net/http"
	"time"

	"github.com/kindiregg/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
