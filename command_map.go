package main

import (
	"encoding/json"
	"fmt"
	"github.com/mdbox037a/pokedexcli/internal/api"
)

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(currentState *state) error {
	url := currentState.next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	response, err := api.GetPokeAPI(url)
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
	url := currentState.previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	response, err := api.GetPokeAPI(url)
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
