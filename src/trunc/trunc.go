// Tauno Erik 24.07.2021
/* Write a program which prompts the user to enter a floating point
number and prints the integer which is a truncated version of the
floating point number that was entered. Truncation is the process of
removing the digits to the right of the decimal place. */
package main

import (
	"fmt"     // format
	"strconv" // String convert
)

func main() {

	var num float64

	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&num) // Store user input

	var i int = int(num)           // Convert float to int
	var s string = strconv.Itoa(i) // Convert int to string

	fmt.Printf(s + "\n") // Print string and newline
}
