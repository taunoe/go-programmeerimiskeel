# Interfaces

* Dynamic type
* Dynamic value

```Go
type Shape2D interface {
  Area() float64
  Perimeter() float64
}

type Triangle{...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimater() float64 {...}
```

* Triangle type satifies the Shape2D interface
* No need to state it explicity

## Defining an interface type

```Go
type Speaker interface {
  Speak ()
}

type Dog struct {
  name string
}

func (d Dog) Speak() {
  fmt.Println(d.name)
}

func main() {
  var s1 Speaker
  var d1 Dog{"Brian"}
  s1 = d1
  s1.Speak()
}
```

* Dynamic type is Dog, dynamic value is d1

## Using interfaces

```Go
type Shape2d interface {
  Area() float64
  Perimeter() float64
}

type Triangle {...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimeter() float64 {...}
type Rectangle {...}
func (t Rectangle) Area() float64 {...}
func (t Rectangle) Perimeter() float64 {...}

func FitInYard (s Shape2d) bool {
  if (s.Area() > 100 && s.Perimeter() > 100) {
    return true
  }
  return false
}

// Type assertions for disambiguation
func DrawShape(s Shape2D) bool {
  rect, ok :=s.(Rectangle)
  if ok {
    DrawRect(rect)
  }
  rect, ok :=s.(Triangle)
  if ok {
    DrawTri(tri)
  }
}

func DrawRect (r Rectangle) {...}
func DrawTri (t Triangle) {...}


// Type switch
func DrawShape(s Shape2D) bool {
  switch sh := s.(type) {
    case Rectangle:
      DrawRect(sh)
    case Triangle:
      DrawTri(sh)
  }
}

```

## Error interface

* Correct operation error==nil

```Go
type error interface {
  Error() string
}
```

## Handling Errors

```Go
f, err := os.Open("test.txt")
if err != nil {
  fmt.Println(err)
  return
}
```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)