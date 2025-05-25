package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ExploreLocation(location_name string) (PokeEncounters, error) {
	url := apiURL + "/location-area/" + location_name

	if val, exists := client.cache.Get(url); exists {
		encountersResp := PokeEncounters{}
		err := json.Unmarshal(val, &encountersResp)
		if err != nil {
			return PokeEncounters{}, err
		}
		return encountersResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeEncounters{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return PokeEncounters{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	fmt.Println(res)
	if err != nil {
		return PokeEncounters{}, err
	}

	encountersResp := PokeEncounters{}
	err = json.Unmarshal(data, &encountersResp)
	if err != nil {
		return PokeEncounters{}, err
	}

	client.cache.Add(url, data)

	return encountersResp, nil
}
