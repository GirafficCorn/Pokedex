package main

import (
	"fmt"
)

func pokedex(c *config, n string) error {
	fmt.Println("Your Pokedex: ")
	for _, x := range c.Pokedex {
		fmt.Printf("  -%s\n", x.Name)
	}
	return nil
}
