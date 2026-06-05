package main

import (
	"encoding/json"
	"fmt"
	"github.com/mdbox037a/pokedexcli/internal/api"
)

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(currentState *state) error {
	var url string
	if currentState.next == nil {
		url = api.BaseURL + api.MapEndpoint
	} else {
		url = *currentState.next
	}
	response, err := api.GetPokeAPI(url, &currentState.dexClient)
	if err != nil {
		return err
	}

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

func commandMapB(currentState *state) error {
	var url string
	if currentState.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = *currentState.previous
	}
	response, err := api.GetPokeAPI(url, &currentState.dexClient)
	if err != nil {
		return err
	}

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
