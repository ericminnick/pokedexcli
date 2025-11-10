package main

import (
	"time"
	"github.com/ericminnick/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)	
	cfg := &configCommand{
		pokedex: make(map[string]Pokemon),
		pokeapiClient: pokeClient,
	}
	replStart(cfg)
}


