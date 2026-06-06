package main

import (
	"fmt"
)

func commandPokedex(currentState *state, args ...string) error {
	if len(currentState.pokedex) == 0 {
		fmt.Println("No pokemon caught yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for p, _ := range currentState.pokedex {
		fmt.Printf(" - %s\n", p)
	}
	return nil
}
