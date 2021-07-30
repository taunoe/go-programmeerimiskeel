# Concurrency in Go

Goroutines - like a tread in Go.

* One goroutine is created automatically to execute the main()
* Other goroutines are created using the **go** keyword

```Go
func main() {

}
```

## sync.WaitGroup

```Go
func foo(wg *sync.WaitGroup) {
  fmt.Printf("New routine")
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  wg.Add(1)
  go foo(&wg)
  wg.Wait()
  fmt.Printf("Main routine")
}
```

## channels

* Transfer data between goroutines
* Channels are typed
* make() to create a channel: c: = make(chan int)
* Send and receive data using the <- operator. Send data on a channel: c <- 3 Receive data from a channel: x := <- c

```Go
func prod(v1 int, v2 int, c chan int) {
  c <- v1*v2
}

func main() {
  c := make(chan int)
  go prod(1, 2, c)
  go prod(3, 4, c)
  a := <- c
  b := <- c
  fmt.Println(a*b)
}
```

### Iterating through a channel

```Go
for i := range c {
  fmt.Println(i)
}
```

### Select statement

```Go
select {
  case a = <- c1:
    fmt.Println(a)
  case b = <- c2:
    fmt.Println(b)
  default:
    fmt.Println("nop")
}

// Select Send or Receive
select {
  case a = <- inchan:
    fmt.Println("Received a")
  case outchan = <- b:
    fmt.Println("Send b")
}

// Select with an Abort Channel
for {
  select {
    case a <- c:
      fmt.Println(a)
    case <-abort:
      return
  }
}
```

## sync.Mutex

Ensures mutual exclusion

* Lock() - method puts the flag up (shared variable in use)
* UnLock - done using shared variable

```Go
var i int = 0
var nut sync.Mutex

func inc() {
  mut.Lock()
  i = i + 1
  mut.Unlock()
}
```

## sync.Once

* Has one method once.Do(f)
* Function f is executed only one time. Even if it is called in multiple goroutines.

```Go
var on sync.Once

func setup() {
  fmt.Println("Init")
}

func dostuff() {
  on.Do(setup)
  fmt.Println("Hello")
  wg.Done
}

var wg sync.WaitGroup

func main() {
  wg.Add(2)
  go dostuff()
  go dostuff()
  wg.Wait
}
```

## Chopsticks and philosophers

```Go
package main

import (
  "fmt"
  "sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
  leftCS, rightCS *ChopS
}

func (p Philo) eat() {
  for {
    p.leftCS.Lock()
    p.rightCS.Lock()

    fmt.Println("eating")

    p.rightCS.Unlock()
    p.leftCS.Unlock()
  }
}

func main() {
	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]} // left and right
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
}

```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)
