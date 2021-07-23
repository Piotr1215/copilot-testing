package main

import (
	"fmt"
)

func main() {
	fmt.Println(twofer("Bob"))
	fmt.Println(twofer("Alice"))
	fmt.Println(twofer(""))
}

// Sample code
func twofer(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}




