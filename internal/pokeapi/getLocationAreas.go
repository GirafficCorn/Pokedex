package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(url string) (Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("response failed with status code %d", res.StatusCode)
	}
	if err != nil {
		return Location{}, err
	}
	var loc Location

	err = json.Unmarshal(body, &loc)
	if err != nil {
		return Location{}, err
	}

	return loc, nil
}
