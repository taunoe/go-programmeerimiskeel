# Map

* Impementation of a hash table

```Go
var idMap map[string]int // Key type, Value type
idMap = make(map[string]int)
```

Map literal:

```Go
idMap := map[string]int {
  "joe": 123
}
```

```Go
// Print "joe" value
fmt.Println(idMap["joe"])

// Lisa/muuda key/value:
idMap["jane"] = 456

// Kustuta
delete(idMap, "joe")
```

```Go
// Two-value assignment tests for existentce of the key
// id is value
// p is presence of key
id, p = idMap["jane"]

// Number of values
fmt.Println(len(idMap))

for key, val := range idMap {
  fmt.Println(key, val)
}
```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)