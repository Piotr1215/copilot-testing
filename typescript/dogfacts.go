package main

// Implement client for this API: https://dog-facts-api.herokuapp.com/api/v1/resources/dogs?number=1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// create type for dog
type dog struct {
	Name string `json:"name"`
	Breed string `json:"breed"`
	Age int `json:"age"`
}

// unmarshal response to json
func unmarshalResponseToJson(body []byte) {
	var dog dog
	err := json.Unmarshal(body, &dog)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dog.Name)
	fmt.Println(dog.Breed)
	fmt.Println(dog.Age)
}

func main() {
	// Create a client
	client := &http.Client{}

	// Create a URL
	u, _ := url.Parse("https://dog-facts-api.herokuapp.com/api/v1/resources/dogs")

	// Create a request
	req, _ := http.NewRequest("GET", u.String(), nil)

	// Send the request
	resp, _ := client.Do(req)

	// Read the response
	body, _ := ioutil.ReadAll(resp.Body)

	unmarshalResponseToJson(body)

	// Print the response
	fmt.Println(string(body))
}





