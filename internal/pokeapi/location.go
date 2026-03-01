package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocation(locationName string) (LocationArea, error) {
	location := LocationArea{}
	fullURL := fmt.Sprintf("%s%s/%s", BASEURL, LOCATIONAREA, locationName)
	data, err := c.searchData(fullURL)
	if err != nil {
		return location, err
	}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return location, err
	}
	return location, nil
}
