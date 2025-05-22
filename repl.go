package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iegpeppino/pokedexcli/api"
)

type Config struct {
	apiClient         api.Client
	NextLocations     *string
	PreviousLocations *string
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

		// Check if the first word of the input
		// matches a command on our map
		command, ok := getCommands()[line[0]]
		// If it matches, makes the callback
		if ok {
			err := command.callback(config)
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
	var words []string
	output := strings.ToLower(text)
	words = strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
	}
}
