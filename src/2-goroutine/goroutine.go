// Tauno Erik 28.07.2021
// User input a series of integers
// Store in array
// Partition the array into 4 parts,
// eash of which is sorted by different goroutine.
// Each goroutine should print the subarray that it will sort.
// The main goroutine should print the entire sorted list.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sort_print(sli []int, i int, wg *sync.WaitGroup) {
	sort.Ints(sli)
	fmt.Printf("Sorted: %d ", i)
	fmt.Println(sli)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Enter integers, separated  with space: ")

	// Init slice
	integers := make([]int, 0, 15)

	// Read input
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	// Remove \n
	input = strings.TrimSuffix(input, "\n")
	// Separate integers
	string_ints := strings.Fields(input)

	for i := 0; i < len(string_ints); i++ {
		// String to int
		in, err := strconv.Atoi(string_ints[i])
		if err != nil {
			fmt.Println(err)
		}
		// add to slice
		integers = append(integers, in)
	}

	size := int((len(integers)) / 4)
	// Four divided subgroups
	tmp0 := make([]int, 0, size)
	tmp1 := make([]int, 0, size)
	tmp2 := make([]int, 0, size)
	tmp3 := make([]int, 0, size)

	// If there are at least four items
	if len(integers) > 3 {
		selector := 0
		for i := 0; i < len(integers); i++ {
			if selector == 4 { // Reset
				selector = 0
			}
			// Add to group
			switch {
			case selector == 0:
				tmp0 = append(tmp0, integers[i])
				selector++
			case selector == 1:
				tmp1 = append(tmp1, integers[i])
				selector++
			case selector == 2:
				tmp2 = append(tmp2, integers[i])
				selector++
			case selector == 3:
				tmp3 = append(tmp3, integers[i])
				selector++
			}
		}
	}

	wg.Add(4)
	go sort_print(tmp0, 0, &wg)
	go sort_print(tmp1, 1, &wg)
	go sort_print(tmp2, 2, &wg)
	go sort_print(tmp3, 3, &wg)
	wg.Wait()

	// Merged slice
	merged := make([]int, 0, len(integers))
	merged = append(merged, tmp0...)
	merged = append(merged, tmp1...)
	merged = append(merged, tmp2...)
	merged = append(merged, tmp3...)

	sort.Ints(merged)
	fmt.Printf("Merged and sorted: ")
	fmt.Println(merged)

}
