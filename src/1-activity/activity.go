// Tauno Erik 26.07.2021
package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
}

/*
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
	return fn
}*/

func main() {
	/* Test
	fn := GenDisplaceFn(10.0, 2.0, 1.0)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
	*/

	// Ask
	var a float64
	fmt.Printf("Enter acceleration: ")
	fmt.Scan(&a)

	var v float64
	fmt.Printf("Enter initial velocity: ")
	fmt.Scan(&v)

	var s float64
	fmt.Printf("Enter initial displacement: ")
	fmt.Scan(&s)

	var t float64
	fmt.Printf("Enter value for time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v, s)

	fmt.Println(fn(t))
}
