package main

import (
	"fmt"
	"os"
)

func commandExit(currentConfig *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
