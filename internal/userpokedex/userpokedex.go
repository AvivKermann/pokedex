package userpokedex

import (
	"errors"
	"sync"

	"github.com/AvivKermann/pokedex/internal/api"
)

type UserPokedex struct {
	mu       sync.Mutex
	Pokemons map[string]api.Pokemon
}

func InitUserPokedex() UserPokedex {
	return UserPokedex{
		Pokemons: make(map[string]api.Pokemon),
	}
}

func (up *UserPokedex) Catch(name string, pokeData api.Pokemon) error {
	up.mu.Lock()
	defer up.mu.Unlock()

	if len(name) <= 0 {
		return errors.New("empty key is not allowed")
	}

	up.Pokemons[name] = pokeData
	return nil

}

func (up *UserPokedex) Get(name string) (api.Pokemon, error) {
	pokemon, exists := up.Pokemons[name]

	if !exists {
		return api.Pokemon{}, errors.New("pokemon hasn't been captured")
	}

	return pokemon, nil

}
