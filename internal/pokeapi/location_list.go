package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// List location areas from the pokeapi
func (c *Client) ListLocations(pageURL *string) (LocationAreasRes, error) {
	url := baseURL + "/location-area"
	// if the provided url is not nil, use that as the request url
	if pageURL != nil {
		url = *pageURL
	}

	// check if the page of location areas is already in the cache
	// if so, read value from the cache
	// if not in the cache, continue to the rest of the function
	if val, ok := c.cache.Get(url); ok {
		locationsRes := LocationAreasRes{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return LocationAreasRes{}, err
		}
		return locationsRes, nil
	}

	// create an HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasRes{}, err
	}

	// do the HTTP request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasRes{}, err
	}
	defer resp.Body.Close()

	// read the data from the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	// make a struct for the response and populate it with the unmarshaled json
	locationsRes := LocationAreasRes{}
	if err = json.Unmarshal(data, &locationsRes); err != nil {
		return LocationAreasRes{}, err
	}

	return locationsRes, nil
}