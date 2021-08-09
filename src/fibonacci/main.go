package main

import "fmt"

func main() {
	var pos int

	fmt.Printf("Enter Fibonacci position: ")
	fmt.Scan(&pos)

	fmt.Println(fib_three(pos))
}
