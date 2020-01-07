package main

import "fmt"

type word string

func (w *word) isEmpty() bool {
	return *w == ""
}

func main() {
	testWord := word("")

	fmt.Println(testWord.isEmpty())
}
