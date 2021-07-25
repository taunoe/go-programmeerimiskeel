# Massiivid Array

Fikseeritud suurusega.

```Go
//Pre-defined with values
var y [5]int = [5]{1,2,3,4,5}

x := [3]int {61, 72, 83}

for i, v range x {
  fmt.Printf("index %d, value %d", i, v)
}

```

## Lõigud Slices

A "window" on an underlying array.

Muudetava suurusega (kuni massiivi suuruseni).

* Pointer - alguspunkt
* Pikkus - elementide arv
* Mahtuvus - Maksimum pikkus

```Go
arr := [...]string {"a", "b", "c", "d", "e", "f", "g"}
s1 := arr[1:3]  // "b", "c"
s2 := arr[2:5]  // "c", "d", "e"

fmt.Printf(len(s1), cap(s1))

// Slice literals
s3 := []int{1, 2, 3}
```

### Make

```Go
s4 = make([]int, 10) // Type, length/capacity

s5 = make([]int, 10, 15) // Type, lenght, capacity
```

## Append

Lisab elemendi lõigu (slice) lõppu.

Kui vaja siis suurendab massiivi suurust.

```Go
s6 = make([]int, 0, 3)  // Lenght is 0
s6 = append(s6, 100)

```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)