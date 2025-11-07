package pokeapi

import (
	"net/http"
	"time"
	"github.com/ericminnick/pokedexcli/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),	
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
