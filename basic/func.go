package main

import (
	"fmt"
	"strconv"
	"strings"
)

func print(funcName, result string) {
	fmt.Printf("Function %-20v: %v\n", funcName, result)
}

func pow(number int, degree uint) int {
	if degree <= 1 {
		return number
	}
	return number * pow(number, degree-1)
}

func concatination(array ...string) string {
	result := ""
	for _, part := range array {
		result += part
	}

	return result
}

func stringConverter(data string, middlewarelist ...func(data string) string) string {
	for _, middleware := range middlewarelist {
		data = middleware(data)
	}
	return data
}

func main() {
	print("print", "some data")
	print("Some concat", concatination("a", "b", "c", "d"))
	print("pow", strconv.Itoa(pow(2, 5)))
	print("String Converter", stringConverter("\n\t SoMe TeXt\n\r\t", strings.ToLower, strings.TrimSpace))
}
