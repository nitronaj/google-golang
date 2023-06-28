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
	"sync"
)

type Chops struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *Chops
}

func (p Philo) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
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
		fmt.Println(left, right)

		philos[i] = &Philo{chopsticks[left], chopsticks[right]}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg)
	}

	wg.Wait()
}
