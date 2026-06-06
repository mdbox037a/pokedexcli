package main

import (
	"errors"
	"fmt"
)

func commandInspect(currentState *state, args ...string) error {
	if len(args) == 0 {
		return errors.New("Error: please supply the name of a caught pokemon to inspect")
	}
	target := args[0]
	if _, caught := currentState.pokedex[target]; !caught {
		fmt.Println("you have not caught that pokemon")
		return nil
		// not returning an error because user cannot query their own pokedex yet
	}

	p := currentState.pokedex[target]
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %s\n", p.Height)
	fmt.Printf("Weight: %s\n", p.Weight)

	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  -%s: %d\n", t.Type.Name)
	}
	return nil
}
