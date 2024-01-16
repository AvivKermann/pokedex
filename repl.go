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
		fields := strings.Fields(inputLower)
		var locName string

		if len(fields) < 1 {
			continue
		}
		inputCommand := fields[0]

		if len(fields) == 2 {
			locName = fields[1]
		}

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
	name        string
	description string
	callback    func(cfg *Config, locName string) error
}

type Config struct {
	next *string
	prev *string
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "help",
			description: "Exits the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displys the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores a certain area in pokemon",
			callback:    commandExplore,
		},
	}

}
