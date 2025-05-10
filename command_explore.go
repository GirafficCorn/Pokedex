package main

import (
	"fmt"

	"github.com/GirafficCorn/Pokedex/internal/pokeapi"
)

func explore(c *config, location string) error {
	if location == "" {
		return nil
	}
	loc, err := pokeapi.GetLocationInformation(location, c.Cache)
	if err != nil {
		return err
	}

	poke_list := loc.PokemonEncounters
	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon: ")
	for _, pok := range poke_list {
		fmt.Printf("- %s\n", pok.Pokemon.Name)
	}
	return nil
}
