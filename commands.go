package main

import (
	"os"	
	"fmt"
	"errors"
	"github.com/ericminnick/pokedexcli/internal/pokeapi"
)


type cliCommand struct {
	name			string
	description		string		
	callback 		func(*configCommand, ...string) error
}

type configCommand struct {
	pokeapiClient	pokeapi.Client
	next			*string
	previous		*string
}


func commandExit(config *configCommand, parameters ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *configCommand, parameters ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for cmd, cmdInfo := range getCommands() {
		fmt.Printf("%v: %v\n", cmd, cmdInfo.description)
	}

	return nil
}

func commandMap(config *configCommand, parameters ...string) error {
	locationResult, err := config.pokeapiClient.ListLocations(config.next)
	if err != nil {
		return err
	}

	config.next = locationResult.Next
	config.previous = locationResult.Previous
	
	for _, loc := range locationResult.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(config *configCommand, parameters ...string) error {
	if config.previous == nil {
		return errors.New("you're on the first page")
	}

	locationResult, err := config.pokeapiClient.ListLocations(config.previous)
	if err != nil {
		return err
	}

	config.next = locationResult.Next
	config.previous = locationResult.Previous
	
	for _, loc := range locationResult.Results {
		fmt.Println(loc.Name)
	}


	return nil
}

func commandExplore(config *configCommand, parameters ...string) error {
	if len(parameters) != 1 {
		return errors.New("you must provide a location name")
	}

	exploreResult, err := config.pokeapiClient.ExploreLocation(parameters[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", exploreResult.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range exploreResult.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
		"help": {
			name: 		"help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:		"map",
			description: "Retrieves and lists map locations 20 at a time",
			callback: commandMap,
		},
		"mapb": {
			name:		"mapb",
			description: "Retrieves and lists previous map locations 20 at a time",
			callback: commandMapb,
		},
		"explore": {
			name:		"explore",
			description: "Retrieves the pokemon encounters, located in specified area (explore <area-name>)",
			callback: commandExplore,
		},
	}
}
