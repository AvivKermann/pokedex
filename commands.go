package main

import(
	"fmt"
	"os"
)

func commandHelp() error {
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

func commandExit()error {
	os.Exit(0)
	return nil
}