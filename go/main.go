package main

// import "github.com/Piotr1215/copilot-testing/co-pilot"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DogFact struct {
	Fact string `json:"fact"`
}

type Client struct {
	BaseURL *url.URL

	httpClient *http.Client
}

// Generate NewClient
func NewClient(baseURL *url.URL) *Client {
	return &Client{
		BaseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func main() {
	req, err := http.NewRequest("GET", "https://dog-facts-api.herokuapp.com/api/v1/resources/dogs?number=2", nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var facts []DogFact

	if err := json.NewDecoder(resp.Body).Decode(&facts); err != nil {
		return nil, err
	}

	return facts, nil
	fmt.Println(facts)

}
