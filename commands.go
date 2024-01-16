package main

import (
	"errors"
	"fmt"
	"os"

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
