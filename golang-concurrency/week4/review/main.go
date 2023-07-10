package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.Mutex
}

type Phil struct {
	id      int
	hostreq *chan int
	hostans *chan int
	leftc   *ChopStick
	rightc  *ChopStick
}

func host(wg *sync.WaitGroup) {

	var num_pmeals int = 0
	var pass_count int = 0
	var psig int = 0 // 0 = can't eat, 1 = can eat, -1 = done eating
	var prc []chan int = make([]chan int, 5)
	var pans []chan int = make([]chan int, 5)
	var pallow []bool
	pallow = make([]bool, 5)

	defer wg.Done()

	// Create send and recieve channels for each philosopher
	for i := 0; i < 5; i++ {
		prc[i] = make(chan int, 25)
		pans[i] = make(chan int, 25)
		pallow[i] = true
	}

	pmeals := make([]int, 5) // number of times philosopher p has eaten
	chop_sticks := make([]*ChopStick, 5)

	for i := 0; i < 5; i++ {
		chop_sticks[i] = new(ChopStick)
	}

	phils := make([]*Phil, 5)
	for i := 0; i < 5; i++ {
		phils[i] = &Phil{i, &prc[i], &pans[i], chop_sticks[i], chop_sticks[(i+1)%5]}
		go phils[i].eat()
		pmeals[i] = 0
	}

	// Main listen loop-- Allowed combinations of dining philosophers: 0 && 2, 0 && 3, 1 && 3, 1 && 4, and 2 && 4. order of each pair does not matter.
	for num_pmeals = 0; num_pmeals <= 15; {

		select {
		case psig = <-prc[0]:
			if pallow[0] && psig == 0 {
				pans[0] <- 1
				pallow[0] = false
				pallow[1] = false
				pallow[4] = false
			} else if !pallow[0] && psig == 0 {
				pans[0] <- 0
			} else if psig == -1 {
				num_pmeals++
				pmeals[0]++
				pallow[0] = true
				pallow[1] = true
				pallow[4] = true
			}

		case psig = <-prc[1]:
			if pallow[1] && psig == 1 {
				pans[1] <- 1
				pallow[1] = false
				pallow[0] = false
				pallow[2] = false
			} else if !pallow[1] && psig == 1 {
				pans[1] <- 0
			} else if psig == -1 {
				num_pmeals++
				pmeals[1]++
				pallow[1] = true
				pallow[0] = true
				pallow[2] = true
			}

		case psig = <-prc[2]:
			if pallow[2] && psig == 2 {
				pans[2] <- 1
				pallow[2] = false
				pallow[1] = false
				pallow[3] = false
			} else if !pallow[2] && psig == 2 {
				pans[2] <- 0
			} else if psig == -1 {
				num_pmeals++
				pmeals[2]++
				pallow[2] = true
				pallow[1] = true
				pallow[3] = true
			}

		case psig = <-prc[3]:
			if pallow[3] && psig == 3 {
				pans[3] <- 1
				pallow[3] = false
				pallow[2] = false
				pallow[4] = false
			} else if !pallow[3] && psig == 3 {
				pans[3] <- 0
			} else if psig == -1 {
				num_pmeals++
				pmeals[3]++
				pallow[3] = true
				pallow[2] = true
				pallow[4] = true
			}

		case psig = <-prc[4]:
			if pallow[4] && psig == 4 {
				pans[4] <- 1
				pallow[4] = false
				pallow[0] = false
				pallow[3] = false
			} else if !pallow[4] && psig == 4 {
				pans[4] <- 0
			} else if psig == -1 {
				num_pmeals++
				pmeals[4]++
				pallow[4] = true
				pallow[0] = true
				pallow[3] = true
			}
		default:
			pass_count++
		} // end of select
		if num_pmeals == 15 {
			break
		}
	}
	for k := 0; k < 5; k++ {
		close(prc[k])
		close(pans[k])
	}
	//fmt.Printf("host:: pass count required for all philosophers to eat %d\n", pass_count)
}

func (p Phil) eat() {

	var eatcount int = 0
	var eatresponse int = 0
	var done_eating = -1

	for eatcount < 3 {

		// send request to eat -- should be philosopher id
		*p.hostreq <- p.id

		eatresponse = <-*p.hostans

		if eatresponse == 1 {
			p.leftc.Lock()
			p.rightc.Lock()
			fmt.Printf("starting to eat number <%d>\n", p.id+1)
			eatcount += 1
			p.leftc.Unlock()
			p.rightc.Unlock()
			fmt.Printf("finishing eating number <%d>\n", p.id+1)
			*p.hostreq <- done_eating
		}

	}

	fmt.Printf("Philosopher %d exiting after eating %d times....\n", p.id+1, eatcount)
}

func main() {
	fmt.Printf("%s\n", "Starting dining philiosphers...")
	var mwg sync.WaitGroup
	fmt.Printf("%s\n", "Launching host go routine...")
	mwg.Add(1)
	go host(&mwg)
	mwg.Wait()
	fmt.Printf("%s\n", "host function complete -- philosphers are done eating...")
}
