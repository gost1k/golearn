package main

import "fmt"

const (
	IotaFirst = 10 << iota
	IotaSecond
	IotaThird
	IotaFourth
	IotaFive
)

func main() {
	fmt.Println(IotaFirst)
	fmt.Println(IotaSecond)
	fmt.Println(IotaThird)
	fmt.Println(IotaFourth)
	fmt.Println(IotaFive)
}
