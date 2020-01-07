package main

import "fmt"

func main() {
	first := 10
	//second := 10
	third := []string{"a", "b", "c", "d", "e", "f"}
	fourth := map[string]string{"a": "A", "b": "B", "c": "C", "d": "D"}

	for {
		fmt.Print(".")
		if first <= 0 {
			break
		}
		first--
	}
	fmt.Println("\n")

	for i := 10; i < 20; i++ {
		fmt.Println("i", i)
	}

	fmt.Println("\n --------------------- \n")

	for i := range third {
		fmt.Println(i, " => ", third[i])
	}

	fmt.Println("\n --------------------- \n")

	for key, value := range fourth {
		fmt.Println(key, " => ", value)
	}

}
