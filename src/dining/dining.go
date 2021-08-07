// Tauno Erik 30.07.2021
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Chopsticks
type ChopS struct {
	sync.Mutex
}

// Philosophers
type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

// Philosopher eat method
func (p Philo) eat() {
	// Eat 3 times
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		//time.Sleep(time.Millisecond * 1)
		fmt.Printf("finishing to eating %d\n", p.number)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Initialization
	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		fmt.Printf("Init CStick %d\n", i)
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	left := rand.Intn(5)
	right := left + 1
	for i := 0; i < 5; i++ {
		fmt.Printf("Init philo %d\n", i+1)
		// Eash philosopher picks up lowest numbered chopstick first
		//philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]} // left and right

		// The philosophers pick up the chopsticks in any order, not lowest-numbered first
		if left == 5 {
			left = 0
			right = 1
		}
		if right == 5 {
			right = 0
		}
		fmt.Printf("Left: %d right: %d\n", left, right)
		philos[i] = &Philo{CSticks[left], CSticks[right], i + 1}
		left++
		right++
	}

	// Dining
	for i := 0; i < 5; i++ {
		go philos[i].eat()
		time.Sleep(time.Millisecond * 2)
	}
}
