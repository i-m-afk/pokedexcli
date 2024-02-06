package user

import (
	"github.com/i-m-afk/pokedexcli/internal/api"
)

type UserPokedex struct {
	inventory map[string]api.Pokemon
}

func InitPokedex() *UserPokedex {
	up := UserPokedex{
		inventory: make(map[string]api.Pokemon),
	}
	return &up
}

func (userPokedex *UserPokedex) AddToPokedex(pokemonInfo api.Pokemon) {
	userPokedex.inventory[pokemonInfo.Name] = pokemonInfo

	// for _, v := range userPokedex.inventory {
	// 	fmt.Println(v.Name)
	// }
}
