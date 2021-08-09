package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(3.00000001))  // 4
	fmt.Println(math.Floor(3.00000001)) // 3
	fmt.Println(math.Min(1, 0))         // 0
	fmt.Println(math.Max(1, 0))         // 1
	fmt.Println(math.Abs(-1))           // 1
	fmt.Println(math.Sqrt(100))         // 10
	fmt.Println(math.Pow(2, 3))         // 8
}
