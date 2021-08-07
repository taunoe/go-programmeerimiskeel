// Tauno Erik 25.07.2021
/* Write a program which prompts the user to enter integers and
stores the integers in a sorted slice. The program should be
written as a loop. Before entering the loop, the program should
create an empty integer slice of size (length) 3. During each
pass through the loop, the program prompts the user to enter an
integer to be added to the slice. The program adds the integer
to the slice, sorts the slice, and prints the contents of the
slice in sorted order. The slice must grow in size to accommodate
any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user
enters the character ‘X’ instead of an integer. */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	// Empty slice
	sli := make([]int, 3)

	// Loop while user enters "x"
	for {
		var input string

		fmt.Printf("Enter a integer: ")
		fmt.Scan(&input) // Store user input:

		if input == "x" {
			break
		}

		// Convert to int
		i, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		// Add to slice:
		sli = append(sli, i)
		// Sort slice
		sort.Ints(sli)
		// Print slice
		fmt.Println(sli)
	}

}
