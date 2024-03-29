package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/AvivKermann/pokedex/internal/api"
)

func commandHelp(cfg *Config, locName string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCliCommands()
	for command := range commands {
		fmt.Printf("%v: %v\n", command, commands[command].description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *Config, locName string) error {
	fmt.Println("Thanks for using my pokedex!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, locName string) error {
	defultURL := "https://pokeapi.co/api/v2/location-area/"

	if cfg.next != nil {
		defultURL = *cfg.next
	}

	resp, err := api.GetLocationAreas(defultURL)

	cfg.next = resp.Next
	cfg.prev = resp.Previous

	if err != nil {
		return err
	}

	for _, name := range resp.Results {
		fmt.Println(name.Name)
	}
	return nil
}

func commandMapb(cfg *Config, locName string) error {
	if cfg.prev == nil {
		return errors.New("already on page one")
	}
	defultURL := *cfg.prev
	resp, err := api.GetLocationAreas(defultURL)

	cfg.next = resp.Next
	cfg.prev = resp.Previous

	if err != nil {
		return err
	}

	for _, name := range resp.Results {
		fmt.Println(name.Name)
	}
	return nil
}

func commandExplore(cfg *Config, locName string) error {
	resp, err := api.GetLocationAreasPokemons(locName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v", locName)
	for _, name := range resp.PokemonEncounters {
		fmt.Println(name.Pokemon.Name)

	}
	return nil

}

func commandCatch(cfg *Config, locName string) error {
	pokemon, baseExperience, err := api.GetPokemonInfo(locName)

	if err != nil {
		return err
	}

	fmt.Printf("Trying to capture %v\n", locName)

	randRange := int(baseExperience / 30)
	userNum, randomNum := rand.Intn(randRange), rand.Intn(randRange)
	time.Sleep(time.Millisecond * 400)

	if userNum == randomNum {
		fmt.Printf("%v Captured\n", locName)
		fmt.Println("You may now inspect it with the inspect command")
		err := userPokedex.Catch(locName, pokemon)
		if err != nil {
			return fmt.Errorf("cannot catch %s error: %v", locName, err)
		}

		return nil
	}
	fmt.Printf("%v, Escaped\n", locName)
	return nil

}
func commandInspact(cfg *Config, locName string) error {
	pokemon, exists := userPokedex.Get(locName)

	if exists != nil {
		return exists
	}
	fmt.Printf("Name: %s\n", locName)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, obj := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", obj.Stat.Name, obj.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf(" -%s\n", types.Type.Name)
	}

	return nil

}
func commandPokedex(cfg *Config, locName string) error {
	if len(userPokedex.Pokemons) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}
	fmt.Println("Your Pokedex: ")
	for _, pokemonName := range userPokedex.Pokemons {
		fmt.Printf("- %s\n", pokemonName.Name)
	}
	return nil
}
