package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *conf) {
	fmt.Println("Welcome to the Pokedex!, type help for usage information")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("pokedex > ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Print(errors.New("Unable to read input"))
		}

		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			continue
		} else {
			fmt.Printf("Unknown command %v\n", commandName)
		}

	}
}

func cleanInput(input string) []string {
	output := input[:len(input)-1]
	output = strings.ToLower(output)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*conf, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of 20 previous location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Explore a location in Pokemon world",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Get information about a captured pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print list of all the names of pokemon the user has caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
