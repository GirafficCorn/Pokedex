package main

import (
	"fmt"
	"os"
)

func exit(c *config, n string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
