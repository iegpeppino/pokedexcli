package main

import (
	"errors"
	"fmt"

	"github.com/iegpeppino/pokedexcli/internal/api"
)

func commandMap(config *Config, pokedex *map[string]api.Pokemon, args ...string) error {
	res, err := config.apiClient.ListLocations(config.NextLocations)
	if err != nil {
		return err
	}

	// Set the new Previous and Next locations URLs
	config.NextLocations = res.Next
	config.PreviousLocations = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(config *Config, pokedex *map[string]api.Pokemon, args ...string) error {
	if config.PreviousLocations == nil {
		return errors.New("there are no previous locations")
	}

	res, err := config.apiClient.ListLocations(config.PreviousLocations)
	if err != nil {
		return err
	}

	// Set the new Previous and Next locations URLs
	config.NextLocations = res.Next
	config.PreviousLocations = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}
