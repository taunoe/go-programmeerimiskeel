# Funktsioonid

* Funktsioon kopeerib sisestatud argumendid/parameetrid.
* Parameetrite muutmine ei muuda neid väljaspool funktsiooni.

```Go
func PrintHello () {
  fmt.Printf("Hello, world.")
}

func foo(x int, y int) {
  fmt.Print(x*y)
}
// or
func foo(x, y int) {
  fmt.Print(x*y)
}

// Return int
func baa(x int) int {
  return x+1
}

// Multiple return values
func baa2(x int) (int, int) {
  return x, x+1
}

func main () {
  PrintHello()
  foo(2, 3)
  y := baa(1)
  a, b := baa2(3)
}
```

## Call by reference

* Kui argumendina kasutada viita (pointer), siis saab funktsioon muuta muutuja väärtust väljaspool funktsiooni.

```Go
func foo (y *int) {
  *y = *y +1
}

func main() {
  x := 2
  foo(&a)
  fmt.Print(x)
}
```

## Passing Array

* Argumenid (massiiv) kopeeritakse

```Go
func foo(x [3]int) int {
  return x[0]
}

func main() {
  a := [3]int{1, 2, 3} // Define array
  fmt.Printf(foo(a)) // Call func
}
```

## Passing array pointers

```Go
func foo(x *[3]int) int {
 (*x)[0] = (*x)[0] + 1
}

func main() {
  a := [3]int{1, 2, 3} // Define array
  foo(&a)  // Call func
  fmt.Print(a) 
}
```

## Pass slices

* Slices contain a pointer to the array

```Go
func foo(sli[] int) int {
  sli[0] = sli[0] + 1
}

func main() {
  a := []int{1, 2, 3} // Define slice
  foo(a)
  fmt.Print(a)
}
```

## Variables as funktsions

```Go
var funcVar func(int) int
  func incFn(x int) int {
  return x+ 1
}

func main() {
  funcVar = incFn
  fmt.Print(funcVar(1))
}
```

## Functions as arguments

```Go
func applyIt(afunct func (int) int, val int) int {
  return afunct(val)
}
```

```Go
func applyIt(afunct func (int) int, val int) int {
  return afunct (val)
}

func incFn(x int) int {
  return x + 1
}

func decFn(x int) int {
  return x - 1
}

func main() {
  fmt.Println(applyIt(incFn, 2))
  fmt.Println(applyIt(decFn, 2))
}
```

## Anonymous functions

```Go
func applyIt(afunct func (int) int, val int) int {
  return afunct (val)
}

func main() {
  v := applyIt(func (x int) int {return x + 1}, 2)
  fmt.Println(v)
}
```

## Functions as return value

```Go
func MakeDistOrigin (o_x, O_y float64)
  func (float64, float64) float64 {
    fn := func (x, y float64) float64 {
      return math.Sqrt(math.Pow(x-o_x, 2) + 
                       math.Pow(y-o_y, 2))
    }
  return fn
}

func main() {
  Dist1 := MakeDistOrigin(0,0)
  Dist2 := MakeDistOrigin(2,2)
  fmt.Println(Dist1(2,2))
  fmt.Println(Dist2(2,2))
}
```

## Variadic function - Variable argument number

Treated as a slice inside function.

```Go
func getMax(vals ...int) int {
  maxV := -1
  for _, v := range vals {
    if v > maxV {
      maxV = v
    }
  }
  return maxV
}

func main() {
  fmt.Println(getMax(1,2,6,4))
  vslice := []int{1,3,6,4}
  fmt.Println(getMax(vslice...))
}
```

## Deferred function calls

* We can delay a function call to the end of the current scope by using the defer keyword
* Typically used for cleanup activities, logging.

```Go
func main() {
  defer fmt.Printlm("Bye!")

  fmt.Println("Hello!")
}
// Hello!
// Bye!
```

```Go
func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!")
  taxRate := .06143
  fmt.Println("Calculating Taxes")
 
  if deductions == 0 || credits == 0 {
    return revenue * taxRate
  }
 
 
  taxValue := (revenue - (deductions * credits)) * taxRate
  if taxValue >= 0 {
    return taxValue
  } else {
    return 0
  }
}
```

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)