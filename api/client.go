package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//take requests and data deserialization,
//and put it all together in our own
// api.Client type.

type Client struct {
	APIKey string
}

func (c *Client) GetTrendingMovies(ctx context.Context) (*Response, error) {
	if c.APIKey == "" {
		return nil, errors.New("API key must be set")
	}
	//must set your own api key 
	requestURL := "https://api.themoviedb.org/3/trending/movie/week?api_key=" + c.APIKey
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}

	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
