package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	pokemonData := Pokemon{}
	fullURL := fmt.Sprintf("%s%s/%s", BASEURL, POKEMON, pokemonName)
	data, err := c.searchData(fullURL)
	if err != nil {
		return pokemonData, err
	}
	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return pokemonData, err
	}
	return pokemonData, nil
}
