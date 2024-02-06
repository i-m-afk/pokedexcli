# POKEDEXCLI

A CLI based pokedex game using [pokeapi](https://pokeapi.co/), written in go.

## Features
- REPL based game
- Implemented caching to reduce api calls
- Go routines to clean old caches and cache size reduce size automatically
- User config store
- User inventory
- Testing (TODO)

## Usage

### Build

`git clone https://github.com/i-m-afk/pokedexcli.git && cd pokedexcli`
`go build && ./pokedexcli`

### Testing

`go test`

### Commands
```
pokedex: Print list of all the names of pokemon the user has caught
exit: Exit the Pokedex
help: Displays the help message
map: Displays the name of 20 location areas in the Pokemon world
mapb: Displays the name of 20 previous location areas in the Pokemon world
explore <location_area>: Explore a location in Pokemon world
catch <pokemon_name>: Catch a pokemon
inspect <pokemon_name>: Get information about a captured pokemon
```
## TODO
- Refactoring
- Usage in-order consistent display
- Use of indexes, to reduce typing
- More testing
