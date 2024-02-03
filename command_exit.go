package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Printf("Exiting...\n")
	os.Exit(0)
	return nil
}
