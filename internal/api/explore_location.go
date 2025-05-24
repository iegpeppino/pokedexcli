package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ExploreLocation(location_name string) (encounters, error) {
	url := apiURL + "/location-area/" + location_name

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
	fmt.Println(res)
	if err != nil {
		return encounters{}, err
	}

	encountersResp := encounters{}
	err = json.Unmarshal(data, &encountersResp)
	if err != nil {
		fmt.Printf("error decoding sakura response: %v\n", err)
		if e, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("syntax error at byte offset %d\n", e.Offset)
		}
		fmt.Printf("sakura response: %q\n", data)
		return encounters{}, err
	}

	client.cache.Add(url, data)

	return encountersResp, nil
}
