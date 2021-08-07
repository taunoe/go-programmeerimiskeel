// Tauno Erik 27.07.2021
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {
		var input string
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n') // ReadString will block until the delimiter is entered
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string
		// Separate animal and info
		tmp := strings.Fields(input)
		ani := tmp[0]
		info := tmp[1]

		switch {
		case ani == "cow":
			if info == "eat" {
				cow.Eat()
			} else if info == "move" {
				cow.Move()
			} else if info == "speak" {
				cow.Speak()
			}
		case ani == "bird":
			if info == "eat" {
				bird.Eat()
			} else if info == "move" {
				bird.Move()
			} else if info == "speak" {
				bird.Speak()
			}
		case ani == "snake":
			if info == "eat" {
				snake.Eat()
			} else if info == "move" {
				snake.Move()
			} else if info == "speak" {
				snake.Speak()
			}
		}
	}

}
