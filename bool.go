package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fouth := !third
	var fifth = true

	fmt.Println("first = ", first)
	fmt.Println("second = ", second)
	fmt.Println("third = ", third)
	fmt.Println("fouth = ", fouth)
	fmt.Println("fifth = ", fifth)

	fmt.Println("2 < 3 is", 2 < 3)
}
