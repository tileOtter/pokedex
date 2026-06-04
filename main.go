package main

import (
	"github.com/tileOtter/pokedex/internal/pokeapi"
)

func main() {
	cfg := pokeapi.Config{}
	startRepl(&cfg)
}
