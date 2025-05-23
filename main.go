package main

import (
	"time"

	"github.com/iegpeppino/pokedexcli/internal/api"
	//"github.com/iegpeppino/pokedexcli/internal/pokecache"
)

func main() {
	// Create a new client with a timeout
	client := api.NewClient(5*time.Second, 3*time.Minute)

	// Initiate config struct for the client
	config := &Config{
		apiClient: client,
	}

	// Create cache
	//cache := pokecache.NewCache(60 * time.Second)

	// Call startRepl using the client's config
	startRepl(config)

}
