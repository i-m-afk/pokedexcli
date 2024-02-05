package main

import "fmt"

func commandHelp(*conf, ...string) error {
	fmt.Printf("Welcome to the Pokedex! \n")
	fmt.Printf("Usage:\n")
	for _, cli := range getCommands() {
		fmt.Printf("%v: %v\n", cli.name, cli.description)
	}
	return nil
}
