package main

import (
	"fmt"
	"os"
	"github.com/AvivKermann/pokedex/internal/api"
	"errors"
)

func commandHelp(cfg *Config) error{
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

func commandExit(cfg *Config) error{
	fmt.Println("Thanks for using my pokedex!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error{
	defultURL := "https://pokeapi.co/api/v2/location/"
	
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

func commandMapb(cfg *Config) error {
	if cfg.prev == nil {
		return errors.New("Cannot go back on page one.")
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