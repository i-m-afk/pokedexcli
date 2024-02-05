package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func commandMapb(config *conf) error {
	locationArea := config.locationArea
	var body []byte
	var err error
	if locationArea.Previous != "" {
		body, err = getLocation(locationArea.Previous, &config.cache)
	} else {
		return errors.New("Previous is Nil")
	}
	if err != nil {
		return errors.New("Some error occured")
	}

	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		log.Fatal("Couldn't unmarshal")
	}
	for _, result := range locationArea.Results {
		fmt.Printf("%s\n", result.Name)
	}
	if locationArea.Next == "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20" &&
		locationArea.Previous == "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20" {
		locationArea.Next = locationArea.Previous
		locationArea.Previous = ""
	}
	config.locationArea = locationArea
	return nil
}
