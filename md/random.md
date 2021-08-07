# Random

```Go
import (
  "math/rand"
  "fmt"
)
 
func main() {
  rand.Seed(time.Now().UnixNano())  //unique number
  fmt.Println(rand.Intn(100))
}
```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)