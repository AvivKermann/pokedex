package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
)

func startRepl() {
	usrInput := bufio.NewScanner(os.Stdin)
	cfg := &Config{}
	for {
		fmt.Printf("pokedex> ")
		usrInput.Scan()
		
		inputLower := strings.ToLower(usrInput.Text())
		inputCommand := strings.Fields(inputLower)[0]
		locName := &inputCommand
		// inputArgs := strings.Fields(inputLower)[1]
		

		command, exist := getCliCommands()[inputCommand]
		if !exist {
			fmt.Println("Unkown command")
			continue
		}
		fmt.Println()
		err := command.callback(cfg, locName)
		if err != nil {
			fmt.Println(err)
		}
		continue
		}
	}

type cliCommand struct {
	
	name string
	description string
	callback func(cfg *Config, locName *string) error
}

type Config struct {
	next *string
	prev *string
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
		"map" : {
			name: "map",
			description : "Displys the next 20 locations",
			callback : commandMap,
		},
		"mapb" : {
			name : "mapb",
			description: "Displays the previous 20 locations",
			callback: commandMapb,
		},
	}

}
