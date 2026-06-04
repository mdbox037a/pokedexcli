package main

import (
	"github.com/mdbox037a/pokedexcli/internal/api"
)

func main() {
	dexClient := api.NewClient(clientTimeout)
	currentState := &state{
		dexClient: dexClient,
	}
	startRepl(currentState)
}
