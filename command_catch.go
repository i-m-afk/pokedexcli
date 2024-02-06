package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/i-m-afk/pokedexcli/internal/api"
)

func commandCatch(config *conf, args ...string) error {
	if len(args) == 0 {
		return errors.New("Usage: catch <pokemon_name>")
	}
	pokemonName := args[0]
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	pokemonInfo := api.Pokemon{}
	data, err := getPokemon(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return errors.New("Pokemon with that name doesn't exists, please check your spelling")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchPokemon := attemptCatch(int32(pokemonInfo.BaseExperience))
	if catchPokemon {
		fmt.Printf("%s was caught\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		config.userPokedex.AddToPokedex(pokemonInfo)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func getPokemon(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func attemptCatch(baseExperience int32) bool {
	catchRate := calcCatchRate(baseExperience, 400)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Float64()
	return randomNumber <= catchRate
}

func calcCatchRate(baseExperience int32, maxBaseExperience int32) float64 {
	baseCatchRate := 0.6
	catchRate := baseCatchRate - (float64(baseExperience) / float64(maxBaseExperience))
	if catchRate < 0 {
		return 0
	}
	return catchRate
}
