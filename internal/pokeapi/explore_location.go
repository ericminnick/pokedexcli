package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(name string) (exploreResponse, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		exploreResp := exploreResponse{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return exploreResponse{}, err
		}
		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil) 
	if err != nil {
		return exploreResponse{}, err 
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return exploreResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return exploreResponse{}, err
	}

	exploreResp := exploreResponse{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return exploreResponse{}, err
	}

	c.cache.Add(url, dat)
	return exploreResp, nil
}
