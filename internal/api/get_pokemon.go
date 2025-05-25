package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(pokemon_name string) (Pokemon, error) {
	url := apiURL + "/pokemon/" + pokemon_name

	// Checks cache for existing data
	if val, exists := client.cache.Get(url); exists {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(url, data)

	return pokemonResp, nil
}
