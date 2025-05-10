package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GirafficCorn/Pokedex/internal/pokecache"
)

func GetPokemon(name string, cache *pokecache.Cache) (Pokemon, error) {
	var body []byte
	url := "https://pokeapi.co/api/v2/pokemon/" + name + "/"
	if r, ok := cache.Get(url); !ok {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		if res.StatusCode == 404 {
			return Pokemon{}, fmt.Errorf("pokemon not found, please check spelling of location")
		}
		if res.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("response failed with status code %d", res.StatusCode)
		}
		cache.Add(url, body)
	} else {
		body = r
	}
	var loc Pokemon

	err := json.Unmarshal(body, &loc)
	if err != nil {
		return Pokemon{}, err
	}

	return loc, nil
}
