package api

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
This client function takes a URL and performs a get request.
If succesful, it returns a list of locations
(pokeLocations struct)
*/
func (client *Client) ListLocations(baseURL *string) (pokeLocations, error) {
	url := apiURL + "/location-area"
	if baseURL != nil {
		url = *baseURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokeLocations{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return pokeLocations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokeLocations{}, err
	}

	locationResponse := pokeLocations{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return pokeLocations{}, err
	}

	return locationResponse, nil
}
