package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	usrInput := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("pokedex> ")
		usrInput.Scan()
		if usrInput.Text() == "exit" {
			commandExit()	
		}
		if usrInput.Text() == "help" {
			commandHelp()
		}

	}
}

type cliCommand struct {
	
	name string
	description string
	callback func() error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help" : {
			name : "help",
			description: "Displays a help message",
			callback : commandHelp,
		},
		"exit" : {
			name : "help",
			description : "Exits the program",
			callback : commandExit,
		},
	}

}


