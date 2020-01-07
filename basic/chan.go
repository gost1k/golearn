package main

import "fmt"

func receive(с chan<- string, prefix string) {
	for data := range с {
		fmt.Println("Read:", prefix, data)
	}
}

func send(с <-chan string, data, prefix string) {
	for _, symbol := range data {
		fmt.Println("Write:", prefix, string(symbol))
		с <- string(symbol)
	}
}

func main() {
	var unBufferChannel chan string
	bufferChannel := make(chan string, 10)

	defer close(unBufferChannel)
	defer close(bufferChannel)

	go receive(unBufferChannel, "pref")
	send(unBufferChannel, "abcdef", "pref")

}
