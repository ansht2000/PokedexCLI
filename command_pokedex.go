package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon, go catch some!")
		return nil
	}

	for name := range cfg.caughtPokemon {
		fmt.Println("  -", name)
	}
	return nil
}