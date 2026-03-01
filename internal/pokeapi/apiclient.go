package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetResource(endpoint string) (Resource, error) {
	resource := Resource{}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return resource, err
	}
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return resource, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return resource, err
	}
	err = json.Unmarshal(data, &resource)
	if err != nil {
		return resource, err
	}
	return resource, nil
}
