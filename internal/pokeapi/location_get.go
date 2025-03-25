package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Get information about pokemon in a location
func (c *Client) GetLocationPok(locationName string) (LocationPokRes, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationsPokRes := LocationPokRes{}
		if err := json.Unmarshal(val, &locationsPokRes); err != nil {
			return LocationPokRes{}, err
		}
		return locationsPokRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokRes{}, err
	}
	
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokRes{}, err
	}
	if res.StatusCode == http.StatusNotFound {
		return LocationPokRes{}, fmt.Errorf("you have provided an invalid area name: %s", locationName)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationPokRes{}, err
	}

	locationPokRes := LocationPokRes{}
	if err = json.Unmarshal(data, &locationPokRes); err != nil {
		return LocationPokRes{}, err
	}

	// value was not found in the cache, add it to the cache
	c.cache.Add(url, data)
	return locationPokRes, nil
}