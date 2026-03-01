package main

import (
	"fmt"
	"math/rand"
)

func capture(pokemonExperience int) bool {
	return rand.Intn(2*pokemonExperience) >= pokemonExperience
}

func commandCatch(c *config, args ...string) error {
	pokemonName := args[0]
	if len(args) == 0 {
		return fmt.Errorf("must provide a pokemon!")
	}
	resp, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("invalid area: %w", err)
	}
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	isCapured := capture(resp.BaseExperience)
	if isCapured {
		fmt.Println(pokemonName + " was caught!")
		c.pokedex[pokemonName] = resp
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}
