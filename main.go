package main

import (
	"log"
	"time"

	"github.com/ansht2000/PokedexCLI/internal/pokeapi"
	"github.com/chzyer/readline"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	rl, err := readline.NewEx(&readline.Config{
		Prompt:      "Pokedex > ",
		HistoryFile: "/tmp/history.txt",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()
	startRepl(cfg, rl)
}
