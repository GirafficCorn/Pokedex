package main

import (
	"fmt"
)

func help(c *config, n string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, com := range getCommands() {
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	return nil
}
