// Tauno Erik 25.07.2021
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	// Create empty slice of struct pointers
	people := []*Person{}

	// Ask filename
	var filename string
	fmt.Printf("Enter a filename: ")
	fmt.Scan(&filename)

	// Open file
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	// Read file content
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	// Iterate through file content
	for _, each_ln := range text {
		// Separate first and last name
		names := strings.Fields(each_ln)

		// Create new person
		tmp := new(Person)
		tmp.fname = names[0]
		tmp.lname = names[1]

		// Add new person to slice
		people = append(people, tmp)
	}

	// Print slice content
	for _, value := range people {
		fmt.Printf("%s %s \n", value.fname, value.lname)
	}

}
