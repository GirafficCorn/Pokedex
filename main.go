package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/GirafficCorn/Pokedex/internal/pokeapi"
	"github.com/GirafficCorn/Pokedex/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		Cache:   pokecache.NewCache(60 * time.Second),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		var param string
		if len(text) > 1 {
			param = text[1]
		} else {
			param = ""
		}

		commands := getCommands()
		command, ok := commands[text[0]]

		if ok {
			err := command.callback(cfg, param)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
