# Märkmed

* Linkedin kursus: [linkedin.com/learning/learning-go-8399317/](https://www.linkedin.com/learning/learning-go-8399317/)
* Harjutus failid [github.com/LinkedInLearning/learning-go-2875237](https://github.com/LinkedInLearning/learning-go-2875237)

```bash
cd hello
go mod init hello
go build .
./hello
go run main.go
```

Built-in functions:

[pkg.go.dev/builtin](https://pkg.go.dev/builtin)

* len(string) - returns yhe lenght of a string
* panic(erroe) - stops execution; display error message
* recover() - manages behaivor of a panicking goroutine

## Variable types

Integers:

* uint8
* int8
* uint16
* int16
* uint32
* int32
* uint 64
* int64

Integer Aliases:

* byte
* uint
* int
* uintptr

Float values:

* float32
* float64

Complex numbers:

* complex64
* complex128

Data collections:

* Arrays
* Slices
* Maps
* Structs

Language organization:

* Functions
* Interfaces
* Channels

Data management:

* Pointers

## Math operators

|name|    |
|----|----|
|sum |   +|
|diffrence| - |
|product | * |
|quotient | / |
|remainder | % |
|bitwise AND | & |
|bitwise OR | \| |
|bitwise XOR| ^ |
|bit clear | &^ |
|left shift | << |
|right shift | >> |

Convert types before using.

```go
var anInt int = 5
var aFloat float64 = 42.1
sum := float(anInt) + aFloat
```

## Build to oter OS

Build to Windows:

```bash
GOOS="windows" go build
```

Build to Linux:

```bash
GOOS="linux" go build
```

Build to macOS:

```bash
GOOS="darwin" go build
```
