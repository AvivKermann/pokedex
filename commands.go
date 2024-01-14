package main

import(
	"fmt"
	"os"
	"github.com/AvivKermann/pokedex/internal/api"
)

func commandHelp() error{
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

func commandExit() error{
	fmt.Println("Thanks for using my pokedex!")
	os.Exit(0)
	return nil
}

func commandMap() error{
	resp, err := api.GetLocationAreas()

	if err != nil {
		return err
	}

	for _, name := range resp.Results {
		fmt.Println(name.Name)
	}
	return nil

	
}