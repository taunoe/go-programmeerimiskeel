package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Strings library
	fmt.Println(strings.ToUpper("hello world"))                        // HELLO WORLD
	fmt.Println(strings.ToLower("HELLO WORLD"))                        // hello world
	fmt.Println(strings.HasPrefix("Hello world", "Hello"))             // true
	fmt.Println(strings.HasSuffix("Hello world", "world"))             // true
	fmt.Println(strings.Replace("Hello world", "Hello", "Goodbye", 1)) // Goodbye world

	// Strings Builder
	var sb strings.Builder
	sb.WriteString("Hello")
	fmt.Println("This is a String Builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Grow(100)

	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())

	y := strconv.Itoa(123) // Int to string
	fmt.Println(y)
	fmt.Printf("%T", y)       // String
	s, err := strconv.Atoi(y) // String to int
	fmt.Println(err)          // nil
	fmt.Printf("%T\n", s)     // int
	fmt.Println(s)

}
