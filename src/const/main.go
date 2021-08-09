package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	const yks int = 1

	fmt.Println(yks)

	// Built in contants
	fmt.Println(math.Pi)       // 3.141592653589793
	fmt.Println(time.December) // December
	fmt.Println(time.UTC)

	const (
		a = 1 // 1
		b     // 1
		c = 2 // 2
		d     // 2
	)
	fmt.Println(a, b, c, d)

	const (
		zero int = iota // 0
		one             // 1
		two             // 2
	)
	fmt.Println(zero, one, two)

	const (
		// binary shift left
		bb = 1 << (10 * iota) // 1
		kb                    // 1024
		mb                    // 1048576
		gb                    // 1073741824
	)
	fmt.Println(bb, kb, mb, gb)

}
