package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// This method looks very similar to the above method
// and could be condensed to one function that both
// methods use, but splitting them up like this
// makes the cfg struct track the current location better
func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're already on the first page")
	}

	locationsRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}