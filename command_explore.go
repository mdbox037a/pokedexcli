package main

import (
	"encoding/json"
	"fmt"
	"github.com/mdbox037a/pokedexcli/internal/api"
)

func commandExplore(currentState *state, location string) error {
	url := api.BaseURL + api.MapEndpoint + "/" + location
	response, err := api.GetPokeAPI(url, &currentState.dexClient)
	if err != nil {
		return err
	}

	//TODO: bookmark june 4 2026
	currentLocations := locationArea{}
	err = json.Unmarshal(response, &currentLocations)
	if err != nil {
		return err
	}
	currentState.previous = currentLocations.Previous
	currentState.next = currentLocations.Next

	for _, location := range currentLocations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
