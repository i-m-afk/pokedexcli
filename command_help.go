package main

import "fmt"

// TODO: sorted order print for consitency

func commandHelp(*conf, ...string) error {
	fmt.Printf("Usage:\n")
	for _, cli := range getCommands() {
		fmt.Printf("%v: %v\n", cli.name, cli.description)
	}
	return nil
}
