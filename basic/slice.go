package main

import "fmt"

func main() {
	first := make([]byte, 5, 15) //Длина и емкость
	second := make([]byte, 3)
	var third []bool
	fourth := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("first	len(%d)	cap(%d)	=	%v\n", len(first), cap(first), first)
	fmt.Printf("second	len(%d)	cap(%d)	=	%v\n", len(second), cap(second), second)
	fmt.Printf("third	len(%d)	cap(%d)	=	%v\n", len(third), cap(third), third)
	fmt.Printf("fourth	len(%d)	cap(%d)	=	%v\n", len(fourth), cap(fourth), fourth)

	fmt.Println("\n --------------------------------- \n")

	part1 := fourth[:2]  // [1 2]
	part2 := fourth[2:4] // [3 4]
	part3 := fourth[4:]  // [5 6]
	part4 := fourth[:]   // [1 2 3 4 5 6]

	fmt.Printf("part1 len(%d) cap(%d)	= %v\n", len(part1), cap(part1), part1)
	fmt.Printf("part2 len(%d) cap(%d)	= %v\n", len(part2), cap(part2), part2)
	fmt.Printf("part3 len(%d) cap(%d)	= %v\n", len(part3), cap(part3), part3)
	fmt.Printf("part4 len(%d) cap(%d)	= %v\n", len(part4), cap(part4), part4)

	fmt.Println("\n --------------------------------- \n")

	part1[0] = 100
	part2[0] = 300
	part3[0] = 500

	fmt.Println("fourth = ", fourth)
	fmt.Println("part4  = ", part4)

	fmt.Println("\n --------------------------------- \n")

	copyOfFourth := make([]int, 6)
	count := copy(copyOfFourth, fourth)
	fourth = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("fourth		= ", fourth)
	fmt.Println("copyOfFourth	= ", copyOfFourth)
	fmt.Println("count			= ", count)

	fmt.Println("\n --------------------------------- \n")

	fourth = append(fourth, copyOfFourth[:3]...)
	fmt.Println("fourth	", fourth)

	fourth = append(fourth, 111, 222, 333)
	fmt.Println("fourth	", fourth)

	fmt.Printf("fourth len(%d) cap(%d)\n", len(fourth), cap(fourth)) //cap 12
	fourth = append(fourth, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("fourth len(%d) cap(%d)\n", len(fourth), cap(fourth)) //cap x2 = 24

}
