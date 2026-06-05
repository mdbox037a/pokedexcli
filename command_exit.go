package main

import (
	"fmt"
	"os"
)

func commandExit(currentState *state, noop_param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
