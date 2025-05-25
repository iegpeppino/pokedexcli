package main

import (
	"fmt"

	"github.com/iegpeppino/pokedexcli/internal/api"
)

func commandHelp(config *Config, pokedex *map[string]api.Pokemon, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s \n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
