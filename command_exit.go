package main

import (
	"fmt"
	"os"

	"github.com/iegpeppino/pokedexcli/internal/api"
)

func commandExit(config *Config, pokedex *map[string]api.Pokemon, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
