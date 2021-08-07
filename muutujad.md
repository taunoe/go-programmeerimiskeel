# Muutujad

Muutuja deklareerimine: `keyword name type`

```Go
var x int  // x = 0
```

Andmetüübi aliased:

```Go
type Celsius float64
```

Muutuja määramine, kasutades andmetüübi aliast:

```Go
var temp Celsius
```

Muutuja algväärtustamine:

```Go
var x int = 100
var y = 100
```

Lühiversioon muutuja määramisest ja algväärtustamisest. Saab kasutada ainult funktsiooni sees.

```Go
x := 100
```

Funktsioon `new()` loob muutuja ja tagastab viida (*pointer*) sellele.

```Go
ptr := new(int)
*ptr = 3
```

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)