package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Work(name string) {
	fmt.Printf("Work: %-5v start\n", name)
	time.Sleep(time.Second)
	fmt.Printf("Work %-5v stop\n", name)
}

//func Sleep() {
//	go Work("simple")
//	time.Sleep(2 * time.Second)
//}

func WaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(i)

		go func(id int) {
			defer wg.Done()
			Work("wg-" + strconv.Itoa(id))
		}(i)
	}

	wg.Wait()
}

func main() {
	fmt.Println("main: start\n")

	WaitGroup()

	fmt.Println("main: stop\n")
}
