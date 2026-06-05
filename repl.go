package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*state, string) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"explore": {
		name:        "explore",
		description: "See pokemon in a location",
		callback:    commandExplore,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Lists nearby locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Lists previous locations",
		callback:    commandMapB,
	},
}

func startRepl(currentState *state) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		if command, ok := commands[input[0]]; ok {
			param1 := ""
			if len(input) == 2 {
				param1 = input[1]
			}
			err := command.callback(currentState, param1)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
