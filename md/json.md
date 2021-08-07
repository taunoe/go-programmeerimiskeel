# JSON Marshalling

```Go
type struct Person {
  Name string
  Addr string
  Phone string
}

p1 := Person(Name:"joe", Addr:"a st.", Phone:"466")

// JSON Marshalling: Go object -> JSON
barr, err := json.Marshal(p1)

// JSON Unmarshalling: JSON -> Go object
var p2 Person
err := json.Unmarshal(barr, &p2)

```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)