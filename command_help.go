package main

import (
	"fmt"
)

// TODO: create dynamic usage entries in helpText
var helpText = `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`

func commandHelp() error {
	fmt.Println(helpText)
	return nil
}
