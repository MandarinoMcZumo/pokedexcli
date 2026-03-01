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
	pokemon, err := c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("invalid area: %w", err)
	}
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	isCapured := capture(pokemon.BaseExperience)
	if isCapured {
		fmt.Println(pokemonName + " was caught!")
		c.pokedex[pokemonName] = pokemon
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}

func commandPokedex(c *config, args ...string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("nothing here...")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key := range c.pokedex {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}

func commandInspect(c *config, args ...string) error {
	pokemonName := args[0]
	if len(args) == 0 {
		return fmt.Errorf("must provide a pokemon!")
	}

	pokemon, ok := c.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf(`
Name: %s
Height: %d
Weight: %d
`, pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, tp := range pokemon.Types {
		fmt.Printf("  - %s\n", tp.Type.Name)
	}
	return nil
}
