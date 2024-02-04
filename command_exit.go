package main

import (
	"fmt"
	"os"
)

func commandExit(config *conf) error {
	fmt.Printf("Exiting...\n")
	os.Exit(0)
	return nil
}
