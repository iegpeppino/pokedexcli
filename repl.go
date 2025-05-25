package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iegpeppino/pokedexcli/internal/api"
)

type Config struct {
	apiClient         api.Client
	NextLocations     *string
	PreviousLocations *string
	Pokedex           map[string]api.Pokemon
}

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		line := cleanInput(scanner.Text())
		// If the command is empty the loop restarts
		if len(line) == 0 {
			continue
		}

		// Logic to catch arguments that follow a command
		// If our clean line slice has more than one word
		// the first word is our command and the rest
		// are taken as arguments
		commandName := line[0]
		args := []string{}
		if len(line) > 1 {
			args = line[1:]
		}

		// Check if the first word of the input
		// matches a command on our map
		command, ok := getCommands()[commandName]
		// If it matches, makes the callback
		if ok {

			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
			// if there's no match it prints "unk command"
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Lists Pokemon located in current area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Throws a pokeball to a pokemon in an attempt to catch it",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Returns stats of a given pokemon if it has been caught",
			callback:    commandInspect,
		},
	}
}
