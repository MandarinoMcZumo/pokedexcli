package main

import (
	"fmt"

	"github.com/mandarinomczumo/pokedexcli/internal/pokeapi"
)

const (
	RESOURCE string = "location-area"
	BASEURL  string = "https://pokeapi.co/api/v2/"
)

func commandMap(c *config) error {
	endpoint := BASEURL + RESOURCE
	if c.Next != "" {
		endpoint = c.Next
	}
	areas, err := pokeapi.GetResource(endpoint)
	if err != nil {
		return fmt.Errorf("invalid resource: %w", err)
	}
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	c.Next = areas.Next
	c.Previous = areas.Previous
	return nil
}

func commandMapBack(c *config) error {
	if c.Previous == "" {
		fmt.Println("can't go back!")
		return nil
	}
	areas, err := pokeapi.GetResource(c.Previous)
	if err != nil {
		return fmt.Errorf("invalid next endpoint: %w", err)
	}
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	c.Next = areas.Next
	c.Previous = areas.Previous
	return nil
}
