package main

import (
	"github.com/mdbox037a/pokedexcli/internal/api"
	"time"
)

const (
	clientTimeout = (5 * time.Second)
)

type state struct {
	dexClient api.Client
	next      string
	previous  string
}
