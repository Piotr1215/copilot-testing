package gocopilot

import (
	"fmt"
)

// Sample code
func twofer(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}




