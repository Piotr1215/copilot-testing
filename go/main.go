package main

// import "github.com/Piotr1215/copilot-testing/co-pilot"

func main() {
	client = NewClient("https://dog-facts-api.herokuapp.com/")
	facts = client.GetTwoRandomDogFactsRequest()
	Println(facts)

}
