package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please, provide the name of a location")
	}

	name := args[0]
	res, err := config.apiClient.ExploreLocation(name)
	if err != nil {
		fmt.Println("the error was here")
		return err
	}

	fmt.Printf("Exploring %s...\n", res.Name)
	fmt.Println("Found Pokemon")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
