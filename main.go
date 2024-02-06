package main

// TODO: catch pokemon by index_number/id and explore map via that

import (
	"github.com/i-m-afk/pokedexcli/internal/api"
	"github.com/i-m-afk/pokedexcli/internal/pokecache"
	"github.com/i-m-afk/pokedexcli/internal/user"
	"time"
)

type conf struct {
	locationArea api.LocationArea
	cache        pokecache.Cache
	userPokedex  user.UserPokedex
}

func main() {
	cache := pokecache.NewCache(5 * time.Second) // refresh cache after x seconds
	config := &conf{
		locationArea: api.LocationArea{},
		cache:        *cache,
		userPokedex:  *user.InitPokedex(),
	}
	startRepl(config)
}
