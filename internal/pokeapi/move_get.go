package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetMovePok(moveName string) (MovesRes, error) {
	url := baseURL + "/move/" + moveName

	if val, ok := c.cache.Get(url); ok {
		movesRes := MovesRes{}
		if err := json.Unmarshal(val, &movesRes); err != nil {
			return MovesRes{}, err
		}
		return movesRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MovesRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return MovesRes{}, err
	}
	if res.StatusCode == http.StatusNotFound {
		return MovesRes{}, fmt.Errorf("you have provided an invalid move name: %s", moveName)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return MovesRes{}, nil
	}

	movesRes := MovesRes{}
	if err = json.Unmarshal(data, &movesRes); err != nil {
		return MovesRes{}, err
	}

	// value was not found in cache, add it to cache here
	c.cache.Add(url, data)
	return movesRes, nil
}