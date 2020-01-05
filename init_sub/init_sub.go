package c

import "fmt"

func init() {
	fmt.Println("INIT_SUB INIT")
}

func Calc(a int, b int) int {
	return a * b
}
