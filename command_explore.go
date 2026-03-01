package main

import (
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	area := args[0]
	if len(args) == 0 {
		return fmt.Errorf("must provide a location!")
	}
	areas, err := c.pokeapiClient.GetLocation(area)
	if err != nil {
		return fmt.Errorf("invalid area: %w", err)
	}
	pokemons := areas.PokemonEncounters
	fmt.Println("Exploring " + area + "...")
	if len(pokemons) == 0 {
		fmt.Println("No Pokemon found!")
		return nil
	}
	fmt.Println("Found Pokemon: ")
	for _, poke := range pokemons {
		fmt.Println("- " + poke.Pokemon.Name)
	}
	return nil
}
