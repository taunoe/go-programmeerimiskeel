package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readFromStdin() []rune {
	stdin := bufio.NewReader(os.Stdin)
	var line []rune
	for {
		c, _, err := stdin.ReadRune()
		if err == io.EOF || c == '\n' {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading standard input\n")
			os.Exit(1)
		}
		line = append(line, c)
	}
	return line
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 { return a*(t*t)/2 + v0*t + s0 }
}

func main() {
	fmt.Println("Please enter acceleration, initial velocity, initial displacement separated by commas:")
	input_str := strings.Split(string(readFromStdin()), ",")
	acceleration, _ := strconv.ParseFloat(input_str[0], 64)
	initial_velocity, _ := strconv.ParseFloat(input_str[1], 64)
	initial_displacement, _ := strconv.ParseFloat(input_str[2], 64)
	fmt.Println("Please a value for time:")
	t, _ := strconv.ParseFloat(string(readFromStdin()), 64)
	gen_displace_func := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)
	fmt.Println(gen_displace_func(t))
}
