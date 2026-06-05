package api

import (
	"github.com/mdbox037a/pokedexcli/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

func NewClient(requestTimeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: requestTimeout,
		},
		pokeCache: *pokecache.NewCache(cacheInterval),
	}
}
