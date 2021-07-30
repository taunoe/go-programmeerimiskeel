// Tauno Erik 28.07.2021
/*
Race condition is where the system’s substantive behavior
is dependent on the sequence or timing of other uncontrollable events.
Concurrency is the art of making progress on multiple tasks simultaneously
*/
package main

import (
	"fmt"
	"time"
)

func runner(name string, max int) {
	fmt.Printf("Runner %s starts!\n", name)
	// Make program sleep for 1 ms
	time.Sleep(time.Millisecond * 1)

	for i := 0; i < max; i++ {
		fmt.Printf("Runner %s stepp %d\n", name, i+1)
		// Make program sleep for 1 ms
		time.Sleep(time.Millisecond * 1)
	}

}

func main() {
	// Go program with concurrency
	// Placing the go command infront of the func call creates a goroution.
	// The goroutines ensures that both functions execute simultaneously.
	go runner("Tom", 100)

	runner("Bob", 100)

	fmt.Println("The end!")
}
