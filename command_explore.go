package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mdbox037a/pokedexcli/internal/api"
)

func commandExplore(currentState *state, args ...string) error {
	if len(args) == 0 {
		return errors.New("Error: please provide a location name")
	}
	location := args[0]

	url := api.BaseURL + api.MapEndpoint + location
	response, err := api.GetPokeAPI(url, &currentState.dexClient)
	if err != nil {
		return err
	}

	currentLocation := locationDetail{}
	err = json.Unmarshal(response, &currentLocation)
	if err != nil {
		return err
	}

	pokeList := getPokemonAtLocation(&currentLocation)
	if len(pokeList) == 0 {
		return errors.New("Error parsing data from location")
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, mon := range pokeList {
		fmt.Printf(" - %s\n", mon)
	}

	return nil
}

func getPokemonAtLocation(currentLocation *locationDetail) []string {
	pokeList := make([]string, 0)
	for _, info := range currentLocation.PokemonEncounters {
		pokeList = append(pokeList, info.Pokemon.Name)
	}
	return pokeList
}
