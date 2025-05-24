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

	// Checks if requested entry is already on cache
	if val, exists := client.cache.Get(url); exists {
		locationResp := pokeLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return pokeLocations{}, err
		}
		return locationResp, nil
	}

	// Make new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokeLocations{}, err
	}

	// Send request to Client
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return pokeLocations{}, err
	}

	// Make sure the response is closed at the end
	defer resp.Body.Close()

	// Read response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokeLocations{}, err
	}

	// Decode response data into a pokeLocations struct
	locationResponse := pokeLocations{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return pokeLocations{}, err
	}

	// Add visited entry to cache
	client.cache.Add(url, data)

	return locationResponse, nil
}
