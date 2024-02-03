package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

// func executeCommand(command string) {
// 	switch command {
// 	case "help":
// 		commandHelp()
// 	case "exit":
// 		commandExit()
// 	}
// }
