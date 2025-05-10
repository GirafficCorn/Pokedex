package main

import (
	"fmt"
)

func inspect(c *config, name string) error {
	if name == "" {
		return nil
	}
	res, ok := c.Pokedex[name]
	if !ok {
		return fmt.Errorf("pokemon not caught yet")
	}

	fmt.Printf(`Name: %s
Height: %v
Weight: %v
Stats: 
`, res.Name, res.Height, res.Weight)
	for _, x := range res.Stats {
		fmt.Printf("  -%s: %v\n", x.Stat.Name, x.BaseStat)
	}
	fmt.Println("Types: ")
	for _, x := range res.Types {
		fmt.Printf("  -%s\n", x.Type.Name)
	}

	return nil
}
