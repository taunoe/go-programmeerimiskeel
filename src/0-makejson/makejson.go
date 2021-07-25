// Tauno Erik 25.07.2021
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*type Person struct {
		Name string `json:"name"`
		Addr string `json:"addr"`
	}*/
	type Person struct {
		Name string
		Addr string
	}

	var name string
	fmt.Printf("Enter a name: ")
	reader1 := bufio.NewReader(os.Stdin)
	name, err1 := reader1.ReadString('\n') // ReadString will block until the delimiter is entered
	if err1 != nil {
		fmt.Println(err1)
	}
	name = strings.TrimSuffix(name, "\n") // remove the delimeter from the string

	var addr string
	fmt.Printf("Enter a address: ")
	reader2 := bufio.NewReader(os.Stdin)
	addr, err2 := reader2.ReadString('\n')
	if err2 != nil {
		fmt.Println(err2)
	}
	addr = strings.TrimSuffix(addr, "\n")

	var p1 Person
	p1.Name = name
	p1.Addr = addr

	//fmt.Printf(p1.name)
	//fmt.Printf(p1.addr)
	//fmt.Printf("%v", p1) // Print struct

	jsonData, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(string(jsonData) + "\n")
}
