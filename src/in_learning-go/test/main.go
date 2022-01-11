package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

/* Muutujate deklareerimine väljaspool funktsiooni */
var tere string = "Tere maailm!"
var nr int = 9
var nimi = "Tauno"
var vanus = 12

/* Konstandid */
const synniaasta int = 1900 // Privat

func main() {

	fmt.Println(tere)
	fmt.Printf("Type: %T\n", tere)
	fmt.Println(nr)
	fmt.Printf("Type: %T\n", nr)
	fmt.Println(nimi)
	fmt.Printf("Type: %T\n", nimi)
	fmt.Println(vanus)
	fmt.Printf("Type: %T\n", vanus)

	/* Funktsiooni sees muutujate deklareerimine: */
	loom := "Kass"
	fmt.Println(loom)
	fmt.Printf("Type %T\n", loom)
	aasta := 2022
	fmt.Println(aasta)
	fmt.Printf("Type %T\n", aasta)

	/* Const */
	fmt.Println(synniaasta)
	fmt.Printf("Type %T\n", synniaasta)

	/* User input */
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Sisesta nimi:")
	input, err := reader.ReadString('\n')
	if err != nil { // Kui on error stdin lugemisel
		fmt.Println(err)
	}
	fmt.Println("Sisestasid:", input)

	fmt.Print("Sisesta number:")
	number, err := reader.ReadString('\n')
	if err != nil { // Kui on error stdin lugemisel
		fmt.Println(err)
	}
	/* Convert string to float */
	ujukoma, err := strconv.ParseFloat(strings.TrimSpace(number), 64)
	if err != nil { // Kui on error stdin lugemisel
		fmt.Println(err)
	}
	fmt.Println("Ujukoma arv:", ujukoma)
	fmt.Printf("Type %T\n\n", ujukoma)

	/* math package */
	var pi = 3.14159
	var rounded = math.Round(pi)
	fmt.Printf("Original: %v, Rounded: %v\n", pi, rounded)

	// Deklaring multible variables
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("intSum=", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum1 := f1 + f2 + f3
	floatSum2 := f1 + f2 + f3
	fmt.Println("floatSum=", floatSum1)

	floatSumX := math.Round(floatSum1)
	fmt.Println("math.Round(floatSum1)=", floatSumX)

	floatSumY := math.Round(floatSum2*100) / 100
	fmt.Println("math.Round(floatSum2*100)/100=", floatSumY)

	circleRadius := 15.5
	circumferece := circleRadius * 2 * math.Pi
	fmt.Printf("circumferece %.2f\n", circumferece)

	/* Dates and time */
	n := time.Now()
	fmt.Println("time.Now = ", n)
	t := time.Date(1983, time.December, 27, 20, 50, 0, 0, time.UTC)
	fmt.Println("t = ", t)
	fmt.Println(t.Format(time.ANSIC))
	parsedTime, err := time.Parse(time.ANSIC, "Tue Dec 27 20:50:00 1983")
	if err != nil { // Kui on error stdin lugemisel
		fmt.Println(err)
		panic("This error message!")
	}
	fmt.Printf("parsedTime is %T\n", parsedTime)
}
