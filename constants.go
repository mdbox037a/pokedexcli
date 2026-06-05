package main

import (
	"github.com/mdbox037a/pokedexcli/internal/api"
	"time"
)

const (
	requestTimeout = (5 * time.Second)
	cacheInterval  = (5 * time.Second)
)

type state struct {
	dexClient api.Client
	next      *string
	previous  *string
}

type locationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
