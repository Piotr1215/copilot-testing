package main

// import "github.com/Piotr1215/copilot-testing/co-pilot"

import (
	"fmt"
	"net/url"

	m "github.com/Piotr1215/copilot-testing/go/gocopilot"
)

func main() {
	dogUrl, _ := url.Parse("https://dog-facts-api.herokuapp.com/")
	client := m.NewClient(dogUrl)
	facts, _ := client.GetTwoRandomDogFactsRequest()
	fmt.Println(facts)

}
