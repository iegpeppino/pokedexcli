package main

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/iegpeppino/pokedexcli/internal/api"
)

func commandCatch(config *Config, pokedex *map[string]api.Pokemon, args ...string) error {
	if len(args) != 1 {
		return errors.New("please enter a pokemon's name to catch")
	}

	name := args[0]
	pokemon, err := config.apiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	// Using MewTwo base experience as max threshold
	max_experience := 1500
	difficulty := max(0, min(pokemon.BaseExperience, max_experience))

	success_chance := 100 - difficulty
	roll := rand.IntN(100)

	if roll <= success_chance {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		(*pokedex)[name] = pokemon

		return nil
	}
	fmt.Printf("%v escaped!\n", pokemon.Name)
	return nil

}
