# Viidad / Pointers

Viit on aadress andmetele mälus.

* Operaator `&` tagastab muutuja aadressi.
* Operaator `*` tagastab andmed, mis on aadressil.

Pointers are variables that specifically store addresses.

```Go
package main
import "fmt"
func main() {
  treasure := "Midagi-kedagi."
  // Prindib addressi
  fmt.Println(&treasure)
}
```

## Dereferencing

```Go
package main
import "fmt"
func main() {
  star := "Polaris"

  starAddress := &star

  *starAddress = "Sirius"
  
  fmt.Println("The actual value of star is", star)
}
```

## Changing Values in Different Scopes

```Go
package main
import "fmt"
func addHundred (numPtr *int) {
  *numPtr += 100
}
 
func main() {
  x := 1
  addHundred(&x)
  fmt.Println(x) // Prints 101
}
```

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)