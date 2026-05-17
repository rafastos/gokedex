package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespLocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocationAreas{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}

	locationsResp := RespLocationAreas{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocationAreas{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
