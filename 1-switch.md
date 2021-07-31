# Switch

```Go
func main() {
  x:=7
  switch {
    case x>3:
      fmt.Printf("1")
    case x>5:
      fmt.Printf("2")
    case x==7:
      fmt.Printf("3")
    default: 
      fmt.Printf("4")
  }

  name := "H. J. Simp"
  switch name {
    case "Butch":
      fmt.Println("Head to Robbers Roost.")
    case "Bonnie":
      fmt.Println("Stay put in Joplin.")
    default:
      fmt.Println("Just hide!")
  }
```

## Scoped Short declaration

```Go
  switch season := "summer" ; season {
  case "summer":
    fmt.Println("Go out and enjoy the sun!")
  }
}
```

## Links

* [Conditionals-Codecademy](https://www.codecademy.com/learn/learn-go/modules/learn-go-conditionals/cheatsheet)

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)