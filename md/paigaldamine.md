# Go

## Paigaldamine

### Ubuntu

`sudo apt install golang`

Ja Visual Studio Code laiendus [GO](https://marketplace.visualstudio.com/items?itemName=golang.go).

Paigaldatud Go versiooni nägemiseks: `go version`

## Töökeskkond / Workspace

Töökeskkonna vaikeasukoht: `/home/userName/go`

`GOPATH` -  

Töökeskkonnd on soovituslik, aga mitte kohustuslik, jagada kolmeks alamkataloogiks:

* `/src` - erinevate projektide lähtekood
* `/pkg` - pakettid (teegid)
* `/bin` - käivitus failid

Import path `host/userororg/project/(dir)/package`

## Hello World

Igas projektis peab olema pakett `main`. Millest luuakse käivitusfail.

```Go
package main
import "fmt"  // Standard package: format

func main() {
    fmt.Printf("Hello, world\n")
}
```

![Go Hello World! program.](../img/Ekraanipilt%202021-07-24%2017-25-36.png)

### Print/Println/Printf

```Go
fmt.Print("I", "am", "cool")
// Iamcool
fmt.Println("I", "am", "cool")
// I am cool
name := "Leslie"
age := 34
fmt.Printf("%v is of type %T", name, name)
// Leslie is of type string
```

```Go
package main
import "fmt"
func main() {
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool
  
  stationName = "Union Square"
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
}
```

### Verbs

* %T - prints out the type of the argument.
* %d - interpolate a number into a string! base10
* %f - limit how precise we are by including a value between the % and f like: %.2f.
* %v - value

```Go
  fmt.Printf("%b\n", 25) // 11001
  fmt.Printf("%d\n", 25) // 25
  fmt.Printf("%o\n", 25) // 31
  fmt.Printf("%x\n", 15) // f
  fmt.Printf("%X\n", 30) // 1E
  fmt.Printf("%q\n", 25) // '\x19'
```

## fmt.Scan()

```Go
package main
import "fmt"
 
func main() {
  var name string
  var age int
  fmt.Println("What's your name and age?")
  fmt.Scan(&name, &age)
  fmt.Printf("You entered %v and %d.\n", name, age)
}
```

[fmt dokumentatsioon](https://pkg.go.dev/fmt)

## Constants

```Go
const earthsGravity = 9.80665
```

## Go tööriistad

* `go build` - programmi kompileerimine
* `go doc` - kuvab paketti dokumentasiooni
* `go fmt` - vormindab lähtekoodi ilusaks
* `go get` - pakettide allalaadimine ja paigaldamine
* `go list` - kuvab kõik paigaldatud paketid
* `go run` - kompileerib ja käivitab programi
* `go test` - jooksutab teste: failid, mis lõppevad `_test.go`

Näide: `go doc fmt.Println`

## Packages

[List of Go standard packages](https://pkg.go.dev/std)

```Go
// Alias
import(
    "fmt"
    t "time"
)
```

## Numeric types

```Go
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

## Sprint/Sprintln/Printf

Don’t print strings, but format them instead.

```Go
func main() {
    grade := "100"
    compliment := "Great job!"
    teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)
 
    fmt.Print(teacherSays)
    // Prints: You scored a 100 on the test! Great job!

    correctAns := "A"
    answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)
    fmt.Print(answer) // Prints: And the correct answer is… A!

    template := "I wish I had a %v."
    pet := "puppy"
    wish := fmt.Sprintf(template, pet)
    fmt.Println(wish)
    // Prints: I wish I had a puppy.
}
```

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)
