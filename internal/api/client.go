package api

import (
	"net/http"
	"time"

	"github.com/iegpeppino/pokedexcli/internal/pokecache"
)

// Declare client struct
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// Create a new client
func NewClient(timeout, interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
