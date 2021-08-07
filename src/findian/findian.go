// Tauno Erik 24.07.2021
/* Write a program which prompts the user to enter a string.
The program searches through the entered string for the
characters ‘i’, ‘a’, and ‘n’. The program should print “Found!”
if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter
if the characters are upper-case or lower-case. */
package main

import (
	"fmt" // format
	"strings"
)

func main() {

	var str string

	fmt.Printf("Enter a string: ")
	fmt.Scan(&str)

	str = strings.ToLower(str) // Conver to lower case
	fmt.Printf(str + "\n")

	i := strings.HasPrefix(str, "i") // True or False
	n := strings.HasSuffix(str, "n") // True or False
	a := strings.Contains(str, "a")  // True or False

	// If all True
	if i && n && a {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
