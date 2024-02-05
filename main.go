package main

import (
	"github.com/i-m-afk/pokedexcli/internal/pokecache"
)

type conf struct {
	locationArea LocationArea
	cache        pokecache.Cache
}

func main() {
	cache := pokecache.NewCache()
	config := &conf{
		locationArea: LocationArea{},
		cache:        cache,
	}
	startRepl(config)
}
