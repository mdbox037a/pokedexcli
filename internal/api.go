package main

import (
	"fmt"
	"io"
	"net/http"
)

func getPokeAPI(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Error: %s", res.Status)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
