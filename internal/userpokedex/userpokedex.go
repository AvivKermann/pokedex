package userpokedex

import (
	"errors"
	"sync"

	"github.com/AvivKermann/pokedex/internal/api"
)

type UserPokedex struct {
	mu       sync.Mutex
	pokemons map[string]api.Pokemon
}

func InitUserPokedex() UserPokedex {
	return UserPokedex{
		pokemons: make(map[string]api.Pokemon),
	}
}

func (up *UserPokedex) Catch(name string, pokeData api.Pokemon) error {
	up.mu.Lock()
	defer up.mu.Unlock()

	if len(name) <= 0 {
		return errors.New("empty key is not allowed")
	}

	up.pokemons[name] = pokeData
	return nil

}
