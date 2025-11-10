package main

import (
	"os"	
	"fmt"
	"errors"
	"math/rand"
	"github.com/ericminnick/pokedexcli/internal/pokeapi"
)


type cliCommand struct {
	name			string
	description		string		
	callback 		func(*configCommand, ...string) error
}

type configCommand struct {
	pokedex 		map[string]Pokemon
	pokeapiClient	pokeapi.Client
	next			*string
	previous		*string
}

type Pokemon struct {
	name 			string
	baseExp			int
	height			int
	weight			int
	stats			map[string]int
	types			[]string
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
	for _, cmdInfo := range getCommands() {
		fmt.Printf("%v: %v\n", cmdInfo.name, cmdInfo.description)
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

func commandCatch(config *configCommand, parameters ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", parameters[0])

	catchResult, err := config.pokeapiClient.CatchPokemon(parameters[0])
	if err != nil {
		return err
	}
		
	retrievedStats := make(map[string]int)
	for _, stat := range catchResult.Stats {
		retrievedStats[stat.Stat.Name] = stat.BaseStat
	}

	typeList := []string{}
	for _, poketype := range catchResult.Types {
		typeList = append(typeList, poketype.Type.Name)
	}

	pokemon := Pokemon{
		name:		catchResult.Name,
		baseExp: 	catchResult.BaseExperience,
		height:		catchResult.Height,
		weight:		catchResult.Weight,
		stats:		retrievedStats,
		types:		typeList,
	}

	if _, ok := config.pokedex[pokemon.name]; ok {
		fmt.Printf("%s already caught\n", pokemon.name)
		return nil
	}
	
	catchChance := 35 + (float64(pokemon.baseExp) - 36.0) * ((40.0)/(608.0-36.0)) 
	catchRoll := rand.Intn(100)
	
	if int(catchChance) < catchRoll {
		fmt.Printf("%s was caught!\n", pokemon.name)
		config.pokedex[pokemon.name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.name)
	}
	return nil
}

func commandInspect(config *configCommand, parameters ...string) error {
	if pokemon, ok := config.pokedex[parameters[0]]; ok {
		fmt.Printf("Name: %s\n", pokemon.name)
		fmt.Printf("Height: %v\n", pokemon.height)
		fmt.Printf("Weight: %v\n", pokemon.weight)
		fmt.Printf("Stats:\n")
		for key, value := range pokemon.stats {
			fmt.Printf("  -%s: %v\n", key, value)
		}
		fmt.Printf("Types:\n")
		for _, poketype := range pokemon.types {
			fmt.Printf("  - %s\n", poketype)
		}
	} else {
		fmt.Printf("you have not caught that pokemon")
	} 

	return nil	
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: 		 	"exit",
			description: 	"Exit the Pokedex",
			callback: 		 commandExit,
		},
		"help": {
			name: 			"help",
			description: 	"Displays a help message",
			callback: 		commandHelp,
		},
		"map": {
			name:			"map",
			description: 	"Retrieves and lists map locations 20 at a time",
			callback: 		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description: 	"Retrieves and lists previous map locations 20 at a time",
			callback: 		commandMapb,
		},
		"explore": {
			name:			"explore <area>",
			description: 	"Retrieves the pokemon encounters, located in specified area",
			callback:	 	commandExplore,
		},
		"catch": {
			name: 			"catch <pokemon>",
			description: 	"Adds a pokemon to your pokedex", 
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect <pokemon>",
			description:	"Inpects a pokemon's pokedex entry",
			callback:		commandInspect,
		},
	}
}
