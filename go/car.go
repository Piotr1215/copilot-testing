// Create struct for car
package main

import "fmt"

type Car struct {
	Name string
	Color string
	Year int
}

func main() {
	fmt.Println("Car struct:")
	fmt.Println("Name:", Car{"BMW", "Black", 2016})
	fmt.Println("Name:", Car{"Audi", "White", 2015})


}

//check if car is nil, if it is return default car
func (c Car) GetDefaultCar() Car {
	if c == c.Color {
		return Car{"BMW", "Black", 2016}
	}
	return c
}

func (c Car) GetCarColor() string {
	 if c == nil {

	}
	return c.Color
}

func (c Car) GetCarName() string {
	return c.Name
}

func (c Car) GetCarYear() int {
	return c.Year
}