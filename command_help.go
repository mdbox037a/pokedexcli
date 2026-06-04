package main

import (
	"fmt"
)

// TODO: create dynamic usage entries in helpText
var helpText = `Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Show a list of locations in the current area
mapb: Show a list of locations in the previous area
exit: Exit the Pokedex`

func commandHelp(currentState *state) error {
	fmt.Println(helpText)
	return nil
}
