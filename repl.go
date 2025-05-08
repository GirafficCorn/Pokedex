package main

import (
	"strings"

	"github.com/GirafficCorn/Pokedex/internal/pokecache"
)

type config struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
}

func cleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	words := strings.Fields(cleaned)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}
