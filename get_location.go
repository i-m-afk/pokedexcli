package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getLocation(url, next, prev string) ([]byte, error) {
	var res *http.Response
	var err error
	if prev != "" {
		res, err = http.Get(prev)
	} else if next != "" {
		res, err = http.Get(next)
	} else {
		res, err = http.Get(url)
	}
	if err != nil {
		log.Printf("Error making GET request: %v", err)
		return nil, errors.New("Unable to make GET request")
	}

	body, err := io.ReadAll(res.Body)

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
