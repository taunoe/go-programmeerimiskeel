# Classes

* No Class keyword
* Associating methods with data
* Use dot notation to call the method

```Go
type MyInt int

func (mi MyInt) Double () int { // Funtksioon Double
  return int(mi*2)
}

func main() {
  v := MyInt(3)
  fmt.Println(v.Double())
}
```

## Struct with methods

```Go
// Struct
type Point struct {
  x float64
  y float64
}

func (p Point) DistToOrigin() {
  t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
  return math.Sqrt(t)
}

func main() {
  p1 := Point(3, 4)
  fmt.Println(p1.DistToOrigin())
}
```

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)