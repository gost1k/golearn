package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var defaultFloat float32
	floatVar := 5.5

	fmt.Printf("defaultFloat			= %v\n", defaultFloat)
	fmt.Printf("floatVar 			(%s) = %v\n", reflect.TypeOf(floatVar), floatVar)

	fmt.Println("MaxFloat32			=", math.MaxFloat32)
	fmt.Println("MinFloat32			=", math.SmallestNonzeroFloat32)

	fmt.Println("MaxFloat64			=", math.MaxFloat64)
	fmt.Println("MinFloat64			=", math.SmallestNonzeroFloat64)
}
