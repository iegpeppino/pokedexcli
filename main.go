package main

import (
	"time"

	"github.com/iegpeppino/pokedexcli/api"
)

func main() {
	// Create a new client with a timeout
	client := api.NewClient(5 * time.Second)

	// Initiate config struct for the client
	config := &Config{
		apiClient: client,
	}

	// Call startRepl using the client's config
	startRepl(config)
}
