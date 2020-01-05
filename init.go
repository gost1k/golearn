package main

import (
	"fmt"
	"github.com/gost1k/golearn/init_sub"
)

func init() {
	fmt.Println("init.go INIT")
}

func main() {
	fmt.Println("init.go MAIN")
	fmt.Println(.Calc(2, 3))
}
