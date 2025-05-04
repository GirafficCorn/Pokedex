package main

import (
	"fmt"
)

func helpCallback(commandsMap map[string]cliCommand) func() error {
	return func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		for _, com := range commandsMap {
			fmt.Printf("%v: %v\n", com.name, com.description)
		}
		return nil
	}
}
