package jokefetcher

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const baseJokeURL = "http://joke.loc8u.com:8888/joke?"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type JokeFetcher struct {
	client HTTPClient
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value struct {
		Categories []string `json:"categories"`
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
	} `json:"value"`
}

func NewJokeFetcher(client HTTPClient) *JokeFetcher {
	return &JokeFetcher{client: client}
}

func (f *JokeFetcher) FetchJoke(firstName, lastName, category string) (*JokeResponse, error) {
	query := make(url.Values)
	query.Set("firstName", firstName)
	query.Set("lastName", lastName)
	query.Set("limitTo", category)

	req, err := http.NewRequest("GET", baseJokeURL+query.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}

	var joke JokeResponse
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		return nil, err
	}

	return &joke, nil
}

