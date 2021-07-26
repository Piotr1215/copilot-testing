package gocopilot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Crate a struct based on this JSON:
/*
[
  {
      "fact": "Endal was the first dog to ride on the London Eye (the characteristic ferris wheel in London, England), and was also the first known dog to successfully use a ATM machine."
  },
  {
      "fact": "At the age of 4 weeks, most dogs have developed the majority of their vocalizations."
  }
]
*/
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

// GetTwoRandomDogFactsRequest returns two random facts about dogs
// from the API https://dog-facts-api.herokuapp.com/api/v1/resources/dogs?number=2
func (c *Client) GetTwoRandomDogFactsRequest() ([]DogFact, error) {

	req, err := http.NewRequest("GET", "https://dog-facts-api.herokuapp.com/api/v1/resources/dogs?number=2", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var facts []DogFact

	if err := json.NewDecoder(resp.Body).Decode(&facts); err != nil {
		return nil, err
	}

	fmt.Println(facts)

	return facts, nil
}
