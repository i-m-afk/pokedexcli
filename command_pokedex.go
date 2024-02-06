package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *conf, args ...string) error {
	userPokedex := config.userPokedex.GetUserInventory()
	if len(userPokedex) == 0 {
		return errors.New("You haven't caught a pokemon yet")
	}
	for key := range userPokedex {
		fmt.Printf("- %s\n", key)
	}
	return nil
}
