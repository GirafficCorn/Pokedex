package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/GirafficCorn/Pokedex/internal/pokeapi"
)

func catchPokemon(c *config, name string) error {
	pok, err := pokeapi.GetPokemon(name, c.Cache)
	if name == "" {
		return nil
	}
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pok.Name)

	randomNumber := rand.Float64()
	baseRate := 0.7
	difficultyFactor := float64(pok.BaseExperience) / 600
	catchRate := baseRate * (1 - difficultyFactor)

	if randomNumber < catchRate {
		fmt.Printf("%s was caught!\n", pok.Name)
		c.Pokedex[pok.Name] = pok
	} else {
		fmt.Printf("%s escaped!\n", pok.Name)
	}

	return nil
}
