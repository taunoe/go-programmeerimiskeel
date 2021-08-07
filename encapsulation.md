# Encapsulation

* Public function to allow accescc to hidden data

data code:

```Go
package data

var x int = 1

func PrintX() {
  fmt.Println(x)
}
```

main code:

```Go
package main

import "data"

func main() {
  data.PrintX()
}
```

## Controlling access to struct

* Need InitMe() to assign hidden data fields

```Go
package data

type Point struct {
  x float64
  y float64
}

func (p *Point) InitMe(xn, yn float64) {
  p.x = xn
  p.y = yn
}

func (p *Point) Scale(v float64) {
  p.x = p.x * v
  p.y = p.y * v
}

func (p *Point) PrintMe() {
  fmt.Println(p.x, p.y)
}
```

```Go
package main

func main() {
  var p data.Point
  p.InitMe(3, 4)
  p.Scale(2)
  p.PrintMe()
}
```

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)