package main

import (
	"fmt"
)

func commandHelp(config *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s \n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
