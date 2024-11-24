package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Get information (mainly name and base exp)
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonRes := Pokemon{}
		if err := json.Unmarshal(val, &pokemonRes); err != nil {
			return Pokemon{}, err
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	if res.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("You have provided an invalid pokemon name: %s", pokemonName)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonRes := Pokemon{}
	if err := json.Unmarshal(data, &pokemonRes); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonRes, nil
}