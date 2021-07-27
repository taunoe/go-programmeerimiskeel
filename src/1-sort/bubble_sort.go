// Tauno Erik 26.07.2021
package main

import (
	"fmt"
	"strconv"
)

func swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func bubble_sort(sli []int) {
	len := len(sli)

	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
}

func main() {
	// Init slice:
	numbers := make([]int, 0, 10)

	// Ask numbers from user:
	for i := 0; i < 10; i++ {
		var input string
		fmt.Printf("Enter the %d number of ten: ", i+1)
		fmt.Scan(&input)
		// Convert to int
		in, err := strconv.Atoi(input)
		// If input is not number
		if err != nil {
			i = i - 1
		} else {
			numbers = append(numbers, in)
		}
	}

	// Sort
	bubble_sort(numbers)

	// Print out
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

}
