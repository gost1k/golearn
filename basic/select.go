package main

import (
	"fmt"
	"time"
)

var (
	readFormat  = "RECEIRVE %-8v: %v\n"
	writeFormat = "SEND 	%-8v: %v\n"
)

func writeToChan(c chan string, data, prefix string) {
	for _, symbol := range data {
		fmt.Printf(writeFormat, prefix, string(symbol))
		c <- string(symbol)
	}
}

func main() {
	first := make(chan string)
	second := make(chan string)

	defer close(first)
	defer close(second)

	go writeToChan(first, "ABCDEF", "FIRST")
	go writeToChan(second, "ABCDEFHG", "SECOND")

	for {
		time.Sleep(50 * time.Millisecond)
		select {
		case firstVal := <-first:
			fmt.Printf(readFormat, "FIRST", firstVal)
		case secondVal := <-second:
			fmt.Printf(readFormat, "Second", secondVal)
		default:
			fmt.Println("Default")
			return
		}

	}

}
