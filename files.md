# Failid

## ioutil

Ei sobi hästi suurte failide jaoks. Loeb kogu faili mällu.

```Go
// Read
data, e := ioutil.ReadFile("test.txt")

// Write
dat = "Tere vana kere!"

err := ioutil.WriteFile("outfile.txt", dat, 0777) // 0777 - Premission
```

## Os

```Go
// Read file
f, err := os.Open("test.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()

// Write
f, err := os.Create("outfile.txt")
barr := []byte{1, 2, 3}
nb, err := f.Write(barr)  // Writes []byte
nb, err := f.WriteString("Tere")  // Writes a string
```

___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)