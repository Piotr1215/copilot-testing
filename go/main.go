package main

// import "copilot/co-pilot/gocopilot"

func main() {
	client = NewClient("https://dog-facts-api.herokuapp.com/")
	facts = client.GetTwoRandomDogFactsRequest()
	 Println(facts)

}