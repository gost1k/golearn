package main

import (
	"fmt"
)

func main() {
	first := make(map[bool]float32, 20)
	second := make(map[string]bool)
	third := map[string]uint{"one": 1, "two": 2, "three": 3}
	fourth := map[string]uint{}

	fmt.Println("first		=", first)
	fmt.Println("second	=", second)
	fmt.Println("third		=", third)
	fmt.Println("fourth	=", fourth)

	fmt.Println("\n --------------------------------- \n")

	first[true] = 100
	first[false] = 500

	second["something"] = true
	second["else"] = false

	third["four"] = 4
	third["five"] = 5

	fmt.Println("first  = ", first)
	fmt.Println("second = ", second)
	fmt.Println("third  = ", third)

	delete(third, "three")
	delete(third, "two")

	fmt.Println("third  = ", third)
}
