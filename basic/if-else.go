package main

import (
	"fmt"
	"os"
)

func main() {
	trueVar := true
	falseVar := false

	if trueVar {
		fmt.Println("if trueVar")
	}

	if falseVar {
		fmt.Println("if falseVar")
	} else if trueVar {
		fmt.Println("else if falseVar")
	} else {
		fmt.Println("else falseVar")
	}

	if _, err := os.Stat("unkown file"); err == nil {
		fmt.Println("unknown file exist", err)
	} else if _, err := os.Stat("/home/gost1k/IdeaProjects/go-learn/if-else.go"); err == nil {
		fmt.Println("real file exist", err)
	} else {
		fmt.Println("File not found")
	}

}
