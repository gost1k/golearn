package main

import (
	"fmt"
	c "github.com/gost1k/golearn/init_sub"
	"github.com/gost1k/golearn/some_sub_folder"
)

func init() {
	fmt.Println("INIT INIT")
}

func main() {
	fmt.Println("INIT MAIN")
	fmt.Println(c.Calc(2, 3))
	fmt.Println(ex.Ex(22, 11))
}
