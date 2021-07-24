// Create command line tool to convert grams to kilograms.
package gocopilot

import (
	"fmt"
)

func main() {
	var weight float64
	fmt.Print("Enter weight in grams: ")
	fmt.Scanf("%f", &weight)
	kg := weight / 1000
	fmt.Printf("%.2f grams is %.2f kilograms\n", weight, kg)
}
