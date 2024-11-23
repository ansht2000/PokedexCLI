package pokeapi

import (
	"net/http"
	"time"
)

// HTTP Client
type Client struct {
	httpClient http.Client
}

// Make a new client
func NewClient (timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}