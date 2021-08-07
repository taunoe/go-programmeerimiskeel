# Struct

Näide: Person Struct
Kolm väärtust: nimi, address, telefon

```Go
type Person struct {
  name string
  addr string
  phone string
}

var p1 Person

// Dot notation
p1.name = "joe" // write
x = p1.addr // read

// Initializing, fields to zero
p2 := new(Person)

// Initializing: struct literal
p3 := Person{name:"joe", addr:"a st.", phone:"123"}
```

```Go
 type fahrenheit int
 type celsius int

 var f fahrenheit = 32
 var c celsius = 0

 fmt.Prinln(f, c)

 c = celsius((f-32)*5/9)

 fmt.Prinln(c)
```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)