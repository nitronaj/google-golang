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

/*
+++++++++++++++++++++++++++++++++ Race Condition +++++++++++++++++++++++++++++++++++
A race condition is a bug in concurrent programming that can occur when two or more goroutines access the same shared data at the same time.
This can happen if the goroutines are not synchronized, which means that they are not guaranteed to execute in a specific order.

Suppose two functions that increment an integer by 1 are launched in two separate goroutines and act on the same variable.
If they are executed sequentially, then the initial integer will increase by 1.
However, now that they are in a race condition, this can’t be guaranteed.

Each function for loop task that increment the variable for 10 times:

So, let’s say that Function 1 reads the variable, but before it can increment and write it back, Function 2 also reads the variable.
Both of the functions now increment the variable and write it back, but the variable value will only increase by 1 because both functions initially read the same value.

Race conditions can be difficult to debug because they are not always reproducible. This is because the order in which the goroutines execute is not deterministic.

There are a few ways to avoid race conditions.
One way is to use synchronization primitives, such as mutexes and channels.
These primitives can be used to ensure that only one goroutine can access a shared resource at a time.

*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg, "1")
	go increment(&wg, "2")
	wg.Wait()
	fmt.Print("Final counter is", counter)
}
