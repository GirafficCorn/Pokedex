package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GirafficCorn/Pokedex/internal/pokecache"
)

func GetLocationAreas(url string, cache *pokecache.Cache) (Location, error) {
	var body []byte
	if r, ok := cache.Get(url); !ok {
		res, err := http.Get(url)
		if err != nil {
			return Location{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}
		if res.StatusCode > 299 {
			return Location{}, fmt.Errorf("response failed with status code %d", res.StatusCode)
		}
		fmt.Println("Fetched")
		cache.Add(url, body)
	} else {
		fmt.Println("Cached")
		body = r
	}
	var loc Location

	err := json.Unmarshal(body, &loc)
	if err != nil {
		return Location{}, err
	}

	return loc, nil
}
