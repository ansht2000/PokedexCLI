package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// List location areas from the pokeapi
func (c *Client) ListLocations(pageURL *string) (LocationAreasRes, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasRes{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasRes{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasRes{}, err
	}

	locationsRes := LocationAreasRes{}
	if err = json.Unmarshal(data, &locationsRes); err != nil {
		return LocationAreasRes{}, err
	}

	return locationsRes, nil
}