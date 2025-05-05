package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleaned := strings.ToLower(text)
	words := strings.Fields(cleaned)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
