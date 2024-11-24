package pokeapi

import (
	"net/http"
	"time"
	"github.com/ansht2000/PokedexCLI/internal/pokecache"
)

// HTTP Client
type Client struct {
	cache 		pokecache.Cache
	httpClient 	http.Client
}

// Make a new client
func NewClient (timeout time.Duration, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}