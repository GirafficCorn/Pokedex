package main

import (
	"strings"

	"github.com/GirafficCorn/Pokedex/internal/pokeapi"
	"github.com/GirafficCorn/Pokedex/internal/pokecache"
)

type config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Pokedex  map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	words := strings.Fields(cleaned)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit,
		},
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemon in a location",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    catchPokemon,
		},
		"inspect": {
			name:        "insepct",
			description: "View informatin about a caught Pokemon",
			callback:    inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View Pokedex",
			callback:    pokedex,
		},
	}
}
