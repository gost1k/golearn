package main

import (
	"fmt"
	"time"
)
import "sync"

var wg sync.WaitGroup // 1

func routine(i int) {
	fmt.Printf("start %v sleep\n", i)
	time.Sleep(time.Second * 5)
	fmt.Printf("routine %v finished\n", i)
	defer wg.Done() // 3
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)     // 2
		go routine(i) // *
	}
	wg.Wait() // 4
	fmt.Println("main finished")
}
