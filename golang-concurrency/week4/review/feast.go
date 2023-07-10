package main

import (
	"fmt"
	"sync"
)

type Chopstick struct{ sync.Mutex }

type Person struct {
	id              int
	left_chopstick  *Chopstick
	right_chopstick *Chopstick
	host            chan int
	wg              *sync.WaitGroup
}

func (person *Person) Eat() {
	for i := 0; i < 3; i++ {
		person.host <- 1
		person.left_chopstick.Lock()
		person.right_chopstick.Lock()
		fmt.Printf("starting to eat %d\n", person.id)
		fmt.Printf("finishing eating %d\n", person.id)
		person.left_chopstick.Unlock()
		person.right_chopstick.Unlock()
		<-person.host
	}
	person.wg.Done()
}

func main() {
	var wg sync.WaitGroup
	chopsticks := make([]*Chopstick, 5)
	host := make(chan int, 2)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}
	for i := 0; i < 5; i++ {
		person := Person{id: i + 1, left_chopstick: chopsticks[i], right_chopstick: chopsticks[(i+1)%5], host: host, wg: &wg}
		wg.Add(1)
		go person.Eat()
	}

	wg.Wait()
}
