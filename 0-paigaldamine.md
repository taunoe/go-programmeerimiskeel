# Go

## Paigaldamine

### Ubuntu

    sudo apt install golang

Ja Visual Studio Code laiendus [GO](https://marketplace.visualstudio.com/items?itemName=golang.go).

Paigaldatud Go versiooni nägemiseks:

    go version

## Töökeskkond / Workspace

Töökeskkonna vaikeasukoht:

    /home/userName/go

`GOPATH` -  

Töökeskkonnd on soovituslik, aga mitte kohustuslik, jagada kolmeks alamkataloogiks:

- `/src` - erinevate projektide lähtekood
- `/pkg` - pakettid (teegid)
- `/bin` - käivitus failid

## Hello World

Igas projektis peab olema pakett `main`. Millest luuakse käivitusfail.

```Go
package main
import "fmt"  // Standard package: format

func main() {
    fmt.Printf("Hello, world\n")
}
```

![Go Hello World! program.](img/Ekraanipilt%202021-07-24%2017-25-36.png)

## Go tööriistad

- `go build` - programmi kompileerimine
- `go doc` - kuvab paketti dokumentasiooni
- `go fmt` - vormindab lähtekoodi ilusaks
- `go get` - pakettide allalaadimine ja paigaldamine
- `go list` - kuvab kõik paigaldatud paketid
- `go run` - kompileerib ja käivitab programi
- `go test` - jooksutab teste: failid, mis lõppevad `_test.go`

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)