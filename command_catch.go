package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/mdbox037a/pokedexcli/internal/api"
)

func commandCatch(currentState *state, args ...string) error {
	if len(args) == 0 {
		return errors.New("Error: please specify a pokemon to attempt to catch")
	}
	target := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", target)

	url := api.BaseURL + api.PokemanEndpoint + target
	response, err := api.GetPokeAPI(url, &currentState.dexClient)
	if err != nil {
		return err
	}

	pokemon := pokemon{}
	err = json.Unmarshal(response, &pokemon)
	if err != nil {
		return err
	}

	if rand.Intn(pokemon.BaseExperience) < catchThreshold {
		fmt.Printf("%s was caught!\n", target)
		if _, exists := currentState.pokedex[target]; !exists {
			currentState.pokedex[target] = pokemon
		}
		return nil
	}
	fmt.Printf("%s escaped!\n", target)
	return nil
}
