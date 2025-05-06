package main

import (
	"fmt"

	"github.com/GirafficCorn/Pokedex/internal/pokeapi"
)

func mapCommand(c *config) error {
	var url string
	if c.Next != "" {
		url = c.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	loc, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}
	c.Next = loc.Next
	c.Previous = loc.Previous

	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func mapbCommand(c *config) error {
	var url string
	if c.Previous != "" {
		url = c.Previous
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
		fmt.Println("you're on the first page")
	}

	loc, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}
	c.Next = loc.Next
	c.Previous = loc.Previous

	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}
	return nil
}
