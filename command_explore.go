package main

import (
	"errors"
	"fmt"
)

// List the pokemon in a given area
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you have to provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocationPok(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Printf("Found Pokemon: \n")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
