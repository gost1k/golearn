package main

import "fmt"

func main() {

	switch {
	case 1 < 2:
		fmt.Println("1 < 2")
	case 2 > 6:
		fmt.Println("2 > 6")
	}

	switch suffix := "jpg"; suffix {
	case "png":
		fallthrough
	case "jpg":
		fallthrough
	case "gif":
		fmt.Println("Image")
	default:
		fmt.Println("Unknown type")
	}

	var someObject interface{}

	switch someObject.(type) {
	case int:
		fmt.Println("INT")
	case bool:
		fmt.Println("Bool")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown type")
	}
}
