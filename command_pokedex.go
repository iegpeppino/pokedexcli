package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *Config, args ...string) error {
	if len(config.Pokedex) == 0 {
		return errors.New("you haven't caught any Pokemon yet, go catch them all")
	}
	for _, pokemon := range config.Pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}
