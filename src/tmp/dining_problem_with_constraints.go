// Implement the dining philosopher’s problem with the following constraints/modifications.
//
//	[x] - There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
//	[x] - Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
//	[x] - The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
//	[x] - In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
//	[x] - The host allows no more than 2 philosophers to eat concurrently.
//	[x] - Each philosopher is numbered, 1 through 5.
//	[x] - When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
//	[x] - When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	id         int
	counter    int
	chopStickL *ChopStick
	chopStickR *ChopStick
}

type Host struct {
	executionsPerJob int
}

func (h Host) feed(philosopher *Philosopher, wg *sync.WaitGroup) {
	//fmt.Println("---> feed before lock: ", philosopher.id)
	if RandBool() {
		philosopher.chopStickL.Lock()
		philosopher.chopStickR.Lock()
	} else {
		philosopher.chopStickR.Lock()
		philosopher.chopStickL.Lock()
	}
	//fmt.Println("---> feed after lock: ", philosopher.id)
	inProgressJobs <- *philosopher

	//start := time.Now()
	fmt.Println("starting to eat: ", philosopher.id)
	//time.Sleep(2 * time.Second)
	//finish := time.Now()

	//duration := time.Since(start)
	//fmt.Println("philosopher.id: ", philosopher.id, "start: ", start, "finish: ", finish, "duration: ", duration)

	fmt.Println("finishing eating: ", philosopher.id)


	//fmt.Println("<--- feed before unlock: ", philosopher.id)
	philosopher.chopStickL.Unlock()
	philosopher.chopStickR.Unlock()
	//fmt.Println("<--- feed after unlock: ", philosopher.id)
	//<- inProgressJobs
	doneJobs <- *philosopher
	//fmt.Println("!!!!feed after channel: ", philosopher.id)
	//philosopher.counter++
	//if philosopher.counter < h.executionsPerJob {
	//	availableJobs <- *philosopher
	//} else {
	//	wg.Done()
	//}
	wg.Done()

}

// serve queues of philosophers
func (h Host) serve(wg *sync.WaitGroup, doWG *sync.WaitGroup) {

	for {
		select {
		case doneJob := <-doneJobs:
			fmt.Println("doneJob", doneJob.id)
			<- inProgressJobs
			doneJob.counter++
			if doneJob.counter < h.executionsPerJob {
				availableJobs <- doneJob
				//<- inProgressJobs
				//inProgressJobs <- doneJob
				//fmt.Println("doneJob - next circle: ", " job - ", doneJob.id, " count - ",doneJob.counter)
				//wg.Done()
			} else {
				fmt.Println("doneJob - quit: ", doneJob.id)
				//wg.Done()
			}

		case nextJob := <-availableJobs:
			fmt.Println("availableJobs: ", nextJob.id)
			wg.Add(1)

			go h.feed(&nextJob, wg)

			//nextJob.counter++
			//if nextJob.counter < h.executionsPerJob {
			//	availableJobs <- nextJob
			//}
		case <-abort:
			fmt.Println("abort!!!!!")
			doWG.Done()
			return
			//default:
			//	fmt.Println("----default----")
			//	//wg.Done()
			//	serveWG.Done()
			//	return
		}
	}

}

type ChopStick struct {
	sync.Mutex
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}


var availableJobs = make(chan Philosopher, 5)
var inProgressJobs = make(chan Philosopher, 2)
var doneJobs = make(chan Philosopher)
var abort = make(chan bool)

var wg sync.WaitGroup
var serveWG sync.WaitGroup

func main() {
	var constrL *ChopStick

	philosophers := make([]*Philosopher,5)
	for i := 0; i < 5; i++ {
		constr := ChopStick{}

		job := Philosopher{
			id:         i + 1,
			counter:    0,
			chopStickL: constrL,
			chopStickR: &constr,
		}

		constrL = &constr

		philosophers[i] = &job
	}

	philosophers[0].chopStickL = constrL

	host := Host{executionsPerJob: 3}

	for _, j := range philosophers {
		//availableJobs <- *j
		wg.Add(1)
		go host.feed(j, &wg)
	}
	//for availableJobs <- job

	serveWG.Add(1)
	go host.serve(&wg, &serveWG)

	wg.Wait()

	//time.Sleep(2 * time.Second)

	abort <- true
	serveWG.Wait()
}

