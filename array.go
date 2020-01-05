package main

import "fmt"

func main() {
	var first [2]bool
	second := [...]int{-1, -2, -3}
	third := [10]uint{1, 2, 3, 4, 5, 6}
	var fourth [2][2]float32
	fifth := [3][2]float32{{1}, {2.3, 4.4}}

	fmt.Printf("first 	len(%d) = %v\n", len(first), first)
	fmt.Printf("second 	len(%d) = %v\n", len(second), second)
	fmt.Printf("third 	len(%d) = %v\n", len(third), third)
	fmt.Printf("fourth 	len(%d) = %v\n", len(fourth), fourth)
	fmt.Printf("fifth 	len(%d) = %v\n", len(fifth), fifth)

	fmt.Println("second[2] = ", second[2])

	copyOfArray := third
	copyOfArray[3] = 500
	copyOfArray[4] = 600

	fmt.Println("copyOfArray[3]", copyOfArray[3])
	fmt.Println("copyOfArray[4]", copyOfArray[4])

	fmt.Println("third", third)
}
