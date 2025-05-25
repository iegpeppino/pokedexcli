package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please enter a pokemon to inspect")
	}

	pokename := args[0]
	pokemon, exists := config.Pokedex[pokename]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, item := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", item.Stat.Name, item.BaseStat)
	}
	fmt.Println("Types:")
	for _, item := range pokemon.Types {
		fmt.Printf("  - %v\n", item.Type.Name)
	}
	return nil

}
