package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Next     *string
	Previous *string
}

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func FetchLocations(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)

	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error during get request: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("response failed, status code: %d\nbody: %s", res.StatusCode, body)
	}
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error reading body: %w", err)
	}
	decoded := LocationAreaResponse{}
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error during unmarshal: %w", err)
	}
	return decoded, nil
}
