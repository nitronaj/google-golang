// Implement the dining philosopher's problem with the following constraints/modifications.

// .1 There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// .2 Each philosopher should eat only 3times (not in an infinite loop as we did in lecture)
// .3 The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// 5. The host allows no more than 2philosophers to eat concurrently.
// 6. Each philosopher is numbered, 1through 5.
// 7. When a philosopher starts eating (after it has obtained necessary locks) ti prints "starting to eat «number»" on a line by itself, where <number> is the number of the philosopher.
// 8. When a philosopher finishes eating (before it has released its locks) ti prints "finishing eating «numbers" on a line by itself, where «number> is the number of the philosopher.

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Chops struct{ sync.Mutex }

type Philo struct {
	id, count int

	leftCS, rightCS *Chops
}

func (p Philo) eat(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	// philoId := <-c
	// fmt.Println("PhiloId from channel ", philoId)

	p.leftCS.Lock()
	p.rightCS.Lock()
	// if philoId == p.id {
	// 	p.count += 1
	// }
	fmt.Println("starting to eat ", p.id+1)
	p.leftCS.Unlock()
	p.rightCS.Unlock()

	fmt.Println("finishing eating", p.id+1)
	fmt.Println("")
}

func host(wg *sync.WaitGroup, philos []*Philo) {
	defer wg.Done()
	var c chan int
	var numberOfProcess int
	for {
		numberOfProcess = 0
		philoId := rand.Intn(5)
		// fmt.Println("Chooseing the philo: ", philoId+1)

		if philos[philoId].count >= 3 {
			continue
		}

		philos[philoId].count = philos[philoId].count + 1

		for _, philo := range philos {
			// fmt.Println("Philo", philo.id+1, philo.count)
			numberOfProcess += philo.count
		}

		if numberOfProcess == 15 {
			break
		}

		// fmt.Println("Number of process", numberOfProcess)
		// fmt.Println("-----------------")
		wg.Add(1)

		go philos[philoId].eat(wg, c)

		// c <- philoId
	}

	// close(c)

}

func main() {
	var wg sync.WaitGroup
	chopsticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chops)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		left := i
		right := (i + 1) % 5
		if left > right {
			right, left = left, right
		}
		philos[i] = &Philo{i, 0, chopsticks[left], chopsticks[right]}
	}
	wg.Add(1)
	go host(&wg, philos)
	wg.Wait()
}
