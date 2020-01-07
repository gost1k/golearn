package main

import "fmt"

func main() {
	fmt.Println("Panic test")

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Panic recover", error)
		}
	}()

	panic("Something going wrong")

	fmt.Println("End of main")
}
