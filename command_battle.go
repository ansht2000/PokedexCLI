package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTwoPokemon = errors.New("you can only battle with 2 pokemon")

func commandBattle(cfg *config, args ...string) error {
	if len(args) != 2 {
		return ErrTwoPokemon
	}

	ownedPoks := []OwnedPokemon{}
	for _, arg := range args {
		pok, err := cfg.pokeapiClient.GetPokemon(arg)
		if err != nil {
			return err
		}
		ownedPok := OwnedPokemon{}
		ownedPok.Name = pok.Species.Name
		ownedPok.Moves = make([]Move, 4)
		ownedPok.HP = pok.Stats[0].BaseStat
		ownedPok.Attack = pok.Stats[1].BaseStat
		ownedPok.Defense = pok.Stats[2].BaseStat
		ownedPok.SpecialAttack = pok.Stats[3].BaseStat
		ownedPok.SpecialDefense = pok.Stats[4].BaseStat
		ownedPok.Speed = pok.Stats[5].BaseStat
		for i := range(4) {
			ownedPok.Moves[i].Name = pok.Moves[i].Move.Name
			ownedPok.Moves[i].URL = pok.Moves[i].Move.URL
		}
		ownedPoks = append(ownedPoks, ownedPok)
	}

	var curPokIndex int
	var nextPokIndex int
	if ownedPoks[0].Speed > ownedPoks[1].Speed {
		curPokIndex = 0
		nextPokIndex = 1
	} else {
		curPokIndex = 1
		nextPokIndex = 0
	}
	fmt.Printf("%s is going first!\n", ownedPoks[curPokIndex].Name)

	for (ownedPoks[0].HP > 0 && ownedPoks[1].HP > 0) {
		moveNum := rand.Intn(4)
		moveName := ownedPoks[curPokIndex].Moves[moveNum].Name
		fmt.Printf("%s used %s!\n", ownedPoks[curPokIndex].Name, moveName)
		move, err := cfg.pokeapiClient.GetMovePok(moveName)
		if err != nil {
			return err
		}
		selfAttack := ownedPoks[curPokIndex].Attack
		oppDefense := ownedPoks[nextPokIndex].Defense
		damage := (move.Power * (selfAttack / oppDefense)) / 50
		ownedPoks[nextPokIndex].HP -= damage
		curPokIndex += 1
		nextPokIndex += 1
		curPokIndex %= 2
		nextPokIndex %= 2
	}
	fmt.Printf("%s fainted!\n", ownedPoks[curPokIndex].Name)
	fmt.Printf("%s won!\n", ownedPoks[nextPokIndex].Name)
	return nil
}