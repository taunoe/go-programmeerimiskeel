// Tauno Erik 27.07.2021
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat()   { fmt.Println(c.food) }
func (c Cow) Move()  { fmt.Println(c.locomotion) }
func (c Cow) Speak() { fmt.Println(c.noise) }

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat()   { fmt.Println(b.food) }
func (b Bird) Move()  { fmt.Println(b.locomotion) }
func (b Bird) Speak() { fmt.Println(b.noise) }

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat()   { fmt.Println(s.food) }
func (s Snake) Move()  { fmt.Println(s.locomotion) }
func (s Snake) Speak() { fmt.Println(s.noise) }

func create_new_animal(mapanimal map[string]Animal, name string, animal_type string) {
	switch animal_type {
	case "cow":
		mapanimal[name] = Cow{food: "grass", locomotion: "walk", noise: "moo"}
		fmt.Println("Created new cow: " + name)
	case "bird":
		mapanimal[name] = Bird{food: "worms", locomotion: "fly", noise: "peep"}
		fmt.Println("Created new bird: " + name)
	case "snake":
		mapanimal[name] = Snake{food: "mice", locomotion: "slither", noise: "hss"}
		fmt.Println("Created new snake: " + name)
	}
}

func info_query(name Animal, info string) {
	switch info {
	case "eat":
		name.Eat()
	case "move":
		name.Move()
	case "speak":
		name.Speak()
	}
}

func main() {
	var mapAnimal map[string]Animal
	mapAnimal = make(map[string]Animal)

	for {
		var input string
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n') // ReadString will block until the delimiter is entered
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string
		// Separate commands
		commands := strings.Fields(input)

		if commands[0] == "newanimal" {
			new_name := commands[1]
			new_type := commands[2] // “cow”, “bird”, or “snake”
			create_new_animal(mapAnimal, new_name, new_type)

		} else if commands[0] == "query" {
			animal_name := commands[1]
			animal_info := commands[2] // “eat”, “move”, or “speak”

			var query_name Animal
			query_name = mapAnimal[animal_name]
			info_query(query_name, animal_info)
		} else {
			fmt.Println("Enter: 'newanimal name type' or 'query name info'.")
		}
	}

}
