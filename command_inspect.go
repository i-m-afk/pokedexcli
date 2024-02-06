package main

import (
	"errors"
	"fmt"

	"github.com/i-m-afk/pokedexcli/internal/api"
)

func commandInspect(config *conf, args ...string) error {
	if len(args) == 0 {
		return errors.New("Usage: inspect <pokemon_name>")
	}
	pokemonName := args[0]
	userInventory := config.userPokedex.GetUserInventory()
	pokemon := userInventory[pokemonName]
	pokemon, ok := userInventory[pokemonName]
	if !ok {
		fmt.Printf("you have not caught that pokemon yet\n")
	} else {
		printStats(pokemon)
	}
	return nil
}

func printStats(pokemon api.Pokemon) {
	fmt.Printf("Name:\t%s\nHeight:\t%d\nWeight:\t%d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %-15s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}
}
