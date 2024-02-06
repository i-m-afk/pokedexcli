package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/i-m-afk/pokedexcli/internal/api"
	"io"
	"log"
	"net/http"
)

func commandExplore(config *conf, args ...string) error {
	if len(args) == 0 {
		return errors.New("Usage: explore <location_area>")
	}
	locationArea := args[0]
	url := "https://pokeapi.co/api/v2/location-area/" + locationArea
	fmt.Printf("Exploring %s...\n", locationArea)
	body, err := getLocationAreaInfo(url)
	if err != nil {
		return err
	}

	locationAreaInfo := api.LocationAreaInfo{}
	err = json.Unmarshal(body, &locationAreaInfo)
	if err != nil {
		log.Fatal("Couldn't unmarshal")
	}
	if len(locationAreaInfo.PokemonEncounters) == 0 {
		fmt.Printf("No pokemon ecountered...\n")
		return nil
	}
	fmt.Printf("Found Pokemon:\n")
	for _, result := range locationAreaInfo.PokemonEncounters {
		fmt.Printf("\t- %s\n", result.Pokemon.Name)
	}

	return nil
}

func getLocationAreaInfo(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Cannot parse the response")
	}
	return body, nil
}
