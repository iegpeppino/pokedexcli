package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) Explore_Location(location_name string) (encounters, error) {
	url := apiURL + "/location-area" + location_name

	if val, exists := client.cache.Get(url); exists {
		encountersResp := encounters{}
		err := json.Unmarshal(val, &encountersResp)
		if err != nil {
			return encounters{}, err
		}
		return encountersResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return encounters{}, err
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return encounters{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return encounters{}, err
	}

	encountersResp := encounters{}
	err = json.Unmarshal(data, &encountersResp)
	if err != nil {
		return encounters{}, err
	}

	client.cache.Add(url, data)

	return encountersResp, nil
}
