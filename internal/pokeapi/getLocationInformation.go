package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GirafficCorn/Pokedex/internal/pokecache"
)

func GetLocationInformation(location string, cache *pokecache.Cache) (LocationInformation, error) {
	var body []byte
	url := "https://pokeapi.co/api/v2/location-area/" + location + "/"
	if r, ok := cache.Get(url); !ok {
		res, err := http.Get(url)
		if err != nil {
			return LocationInformation{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationInformation{}, err
		}
		if res.StatusCode == 404 {
			return LocationInformation{}, fmt.Errorf("location not found, please check spelling of location")
		}
		if res.StatusCode > 299 {
			return LocationInformation{}, fmt.Errorf("response failed with status code %d", res.StatusCode)
		}
		cache.Add(url, body)
	} else {
		body = r
	}
	var loc LocationInformation

	err := json.Unmarshal(body, &loc)
	if err != nil {
		return LocationInformation{}, err
	}

	return loc, nil
}
