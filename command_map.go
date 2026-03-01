package main

import (
	"fmt"

	"github.com/mandarinomczumo/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	endpoint := pokeapi.BASEURL + pokeapi.LOCATIONAREA
	if c.LocationAreaNext != "" {
		endpoint = c.LocationAreaNext
	}
	areas, err := c.pokeapiClient.GetResource(endpoint)
	if err != nil {
		return fmt.Errorf("invalid resource: %w", err)
	}
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	c.LocationAreaNext = areas.Next
	c.LocationAreaPrevious = areas.Previous
	return nil
}

func commandMapBack(c *config) error {
	if c.LocationAreaPrevious == "" {
		fmt.Println("can't go back!")
		return nil
	}
	areas, err := c.pokeapiClient.GetResource(c.LocationAreaPrevious)
	if err != nil {
		return fmt.Errorf("invalid next endpoint: %w", err)
	}
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	c.LocationAreaNext = areas.Next
	c.LocationAreaPrevious = areas.Previous
	return nil
}
