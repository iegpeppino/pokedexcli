package main

import (
	"time"

	"github.com/iegpeppino/pokedexcli/internal/api"
)

func main() {
	// Create a new client with a timeout
	client := api.NewClient(5*time.Second, 3*time.Minute)

	// Initiate config struct for the client
	config := &Config{
		apiClient: client,
	}

	// Instance an empty pokedex
	pokedex := map[string]api.Pokemon{}

	// Create cache
	//cache := pokecache.NewCache(60 * time.Second)

	// Call startRepl using the client's config
	// pass the users pokedex as argument
	startRepl(config, &pokedex)

}
