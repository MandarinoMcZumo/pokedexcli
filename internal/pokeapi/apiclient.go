package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mandarinomczumo/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval * time.Second),
	}
}

func (c *Client) searchData(endpoint string) ([]byte, error) {
	cacheHit, ok := c.cache.Get(endpoint)
	if ok {
		return cacheHit, nil
	}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	err = c.cache.Add(endpoint, data)
	if err != nil {
		fmt.Println("could not add key to the cache: ", endpoint)
	}
	return data, nil
}

func (c *Client) GetUnnamedResource(endpoint string) (Resource, error) {
	resource := Resource{}
	data, err := c.searchData(endpoint)
	if err != nil {
		return resource, err
	}
	err = json.Unmarshal(data, &resource)
	if err != nil {
		return resource, err
	}
	return resource, nil
}
