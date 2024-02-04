package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func commandMap(config *conf) error {
	locationArea := config.locationArea
	var body []byte
	var err error
	if locationArea.Next == "" {
		body, err = getLocation("https://pokeapi.co/api/v2/location-area/", "", "")
	} else {
		body, err = getLocation("", locationArea.Next, "")
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
	// store the current cofig
	config.locationArea = locationArea
	return nil
}