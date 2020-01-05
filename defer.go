package main

import "fmt"

func main() {
	//a := "Hello"
	fmt.Println("MAIN IN")

	defer func() {
		fmt.Println("DEFER IN MAIN")
	}()

	first()

	fmt.Println("MAIN OUT")

}

func first() {
	fmt.Println("FIRST IN")

	defer func() {
		fmt.Println("DEFER IN FIRST")
	}()

	fmt.Println("FIRST OUT")

}
