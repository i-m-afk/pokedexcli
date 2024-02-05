package main

import (
	"errors"
	"fmt"
	"github.com/i-m-afk/pokedexcli/internal/pokecache"
	"io"
	"log"
	"net/http"
)

func getLocation(url string, cache *pokecache.Cache) ([]byte, error) {
	var res *http.Response
	var err error
	body, err := cache.GetDataFromCache(url)
	if err == nil {
		fmt.Println("Cache hit")
		return body, nil // cached data exists
	}
	// new request cache doesn't exists
	res, err = http.Get(url)
	if err != nil {
		log.Printf("Error making GET request: %v", err)
		return nil, errors.New("Unable to make GET request")
	}

	body, err = io.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Cannot parse the response")
	}

	fmt.Println("Cache miss")
	cache.AddDataToCache(url, body)
	return body, nil
}
