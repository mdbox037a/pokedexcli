package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetPokeAPI(url string, dexClient *Client) ([]byte, error) {
	if data, exists := dexClient.pokeCache.Get(url); exists {
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := dexClient.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Request error: %s", res.Status)
	}
	if err != nil {
		return nil, err
	}

	dexClient.pokeCache.Add(url, body)
	return body, nil
}
