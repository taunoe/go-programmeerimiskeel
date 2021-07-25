package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
type Person struct{
   fname string
   lname string
}

func main(){


	var f string

	fmt.Println("Enter Filename with correct path:")

	fmt.Scan(&f)
	//f= "names.txt"

	file, err := os.Open(f)
    if err != nil {
		fmt.Println("Input File Open error")
         return 
    }
    defer file.Close()

	s:=make([]Person,0)
    scanner := bufio.NewScanner(file)
   
	var line string
	var dindex int
	var fn, ln string
    for scanner.Scan() {
			line =scanner.Text()
					//fmt.Println(line)
		
		 	dindex =strings.Index(line," ")
			if dindex>-1 {
				fn = line[:dindex]
				ln= line[dindex:len(line)]
				//fmt.Println(fn,ln)
				p:=Person{fname:fn,lname:ln}
				s = append(s,p)
			}
    }

    if err := scanner.Err(); err != nil {
		fmt.Println("Input File Read error")
    }

	for _, p := range s {

		fmt.Printf("%s, %s\n",p.fname,p.lname )
	}
}