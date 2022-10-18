package namefetcher

import (
	"encoding/json"
	"net/http"
)

const NameURL = "https://names.mcquay.me/api/v0/"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type NameFetcher struct {
	client HTTPClient
}

type NameResponse struct {
	First string `json:"first_name"`
	Last  string `json:"last_name"`
}

func NewNameFetcher(client HTTPClient) *NameFetcher {
	return &NameFetcher{client: client}
}

func (f *NameFetcher) FetchName() (*NameResponse, error) {
	req, err := http.NewRequest("GET", NameURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}

	var name NameResponse
	err = json.NewDecoder(resp.Body).Decode(&name)
	if err != nil {
		return nil, err
	}

	return &name, nil
}
