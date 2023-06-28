package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func increment(wg *sync.WaitGroup, name string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		counter = counter + 1
		fmt.Printf("routine%s counter is: %d\n", name, counter)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg, "1")
	go increment(&wg, "2")
	wg.Wait()
	fmt.Print("Final counter is", counter)
}
