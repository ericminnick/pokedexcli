package pokeapi

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (catchResponse, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		catchResp := catchResponse{}
		err := json.Unmarshal(val, &catchResp)
		if err != nil {
			return catchResponse{}, err
		}
		return catchResp, nil
	}

	req, err := http.NewRequest("GET", url, nil) 
	if err != nil {
		return catchResponse{}, err 
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return catchResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return catchResponse{}, err
	}
	if string(dat) == "Not Found" {
		return catchResponse{}, fmt.Errorf("Not a valid pokemon")
	}
	catchResp := catchResponse{}
	err = json.Unmarshal(dat, &catchResp)
	if err != nil {
		return catchResponse{}, err
	}

	c.cache.Add(url, dat)
	return catchResp, nil
}
