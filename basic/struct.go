package main

import "fmt"

type Car struct {
	name string
}

func (c *Car) getName() {
	fmt.Println("Car", c.name)
}

type Parts struct {
	Car
	Details []Car
}

func (p *Parts) getDetails() {
	fmt.Println("Name:", p.Car.name)

	for _, detail := range p.Details {
		detail.getName()
	}
}

func main() {
	// Simple way
	BMW := Car{
		name: "BMW",
	}
	BMW.getName()

	// Struct with slice
	audi := Parts{
		Details: []Car{{name: "Part1"}, {name: "Part2"}},
	}
	audi.Car.name = "Audi auto"
	audi.getDetails()

}
