package main

import (
	"github.com/i-m-afk/pokedexcli/internal/pokecache"
	"time"
)

type conf struct {
	locationArea LocationArea
	cache        pokecache.Cache
}

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	config := &conf{
		locationArea: LocationArea{},
		cache:        *cache,
	}
	startRepl(config)
}
