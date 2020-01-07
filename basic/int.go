package main

import (
	"fmt"
	"math"
)

func main() {

	var defaultInt int64 // 0
	intVar := 10

	fmt.Println("default", defaultInt)
	fmt.Println("intVar", intVar)

	intVar++
	fmt.Println("IntVar++", intVar)
	intVar--
	fmt.Println("IntVar--", intVar)
	intVar += 10
	fmt.Println("IntVar += 10", intVar)

	intSize()
}

func intSize() {

	fmt.Println("\n")

	fmt.Println("\t Приделы значений целых int \n")

	fmt.Printf("int8:  [%v, %v]\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16: [%v, %v]\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32: [%v, %v]\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64: [%v, %v]\n", math.MinInt64, math.MaxInt64)

	fmt.Println("\t Приделы значений целых uint \n")

	fmt.Printf("uint8:  (0,  %v)\n", math.MaxUint8)
	fmt.Printf("uint16: (0,  %v)\n", math.MaxUint16)
	fmt.Printf("uint32: (0,  %v)\n", math.MaxUint32)
	fmt.Printf("uint64: (0,  %v)\n", ^uint64(0))

}
