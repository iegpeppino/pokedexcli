package main

import (
	"errors"
	"fmt"

	"github.com/iegpeppino/pokedexcli/internal/api"
)

func commandExplore(config *Config, pokedex *map[string]api.Pokemon, args ...string) error {
	if len(args) != 1 {
		return errors.New("please, provide the name of a location")
	}

	name := args[0]
	res, err := config.apiClient.ExploreLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", res.Name)
	fmt.Println("Found Pokemon")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
