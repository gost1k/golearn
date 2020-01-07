package main

import "fmt"

func main() {
	first := "First: line 1 \n line 2"
	second := `Second line \n line 2`

	fmt.Println("first", first)
	fmt.Println("second", second)

	stringOps()
}

func stringOps() {
	first := "first"
	var second string = "second"
	space := " "
	russian := "Русский"
	runeRussian := []rune(russian)

	fmt.Println("first + space + second = ", first+space+second)
	first += space + second
	fmt.Println("first =", first)

	fmt.Println("first[1]				= ", first[1])
	fmt.Println("string(first[1])		= ", string(first[1]))
	fmt.Println("first[3:]				= ", first[3:])
	fmt.Println("first[:7]				= ", first[:7])
	fmt.Println("first[2:7]			= ", first[2:7])

	fmt.Println("rune[3:]				= ", runeRussian[3:])
	fmt.Println("string(rune[3:])		= ", string(runeRussian[3:]))
}
