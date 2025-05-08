package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/GirafficCorn/Pokedex/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		Cache: pokecache.NewCache(10 * time.Second),
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commands := getCommands()
		command, ok := commands[text[0]]
		if ok {
			//Need to init pointer to config
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}
