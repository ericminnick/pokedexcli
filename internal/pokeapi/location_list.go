package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (mapResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.content.Get(url); ok {
		locationsResp := mapResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return mapResponse{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil) 
	if err != nil {
		return mapResponse{}, err 
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return mapResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return mapResponse{}, err
	}

	locationsResp := mapResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return mapResponse{}, err
	}

	c.content.Add(url, dat)
	return locationsResp, nil
}
